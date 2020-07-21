package client

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/entity"
	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/request"
	. "github.com/alec-rabold/EllucianBannerApi-go/pkg/util"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
)

// EllucianAPIClient is the implementation of the API for universities that use Ellucian Banner
type EllucianAPIClient struct {
	collector *colly.Collector
}

// NewEllucianAPIClient creates a new instantiation of the API Service for Ellucian universities
func NewEllucianAPIClient() *EllucianAPIClient {
	return &EllucianAPIClient{
		collector: defaultCollector(),
	}
}

// GetColleges returns a list of all supported colleges to be used for this API
func (e *EllucianAPIClient) GetColleges() []entity.College {
	res := make([]entity.College, 0)
	for key, value := range EllucianSupportedColleges {
		res = append(res, entity.College{
			NameAbbr: key,
			NameFull: value,
		})
	}
	return res
}

// GetTerms returns a list of all academic terms and respective term codes
func (e *EllucianAPIClient) GetTerms(request request.TermsRequestModel) []entity.Term {
	res := make([]entity.Term, 0)
	dom, err := getDocumentModel(request.College, EllucianRegistrationTermsRelativePath, EllucianRegistrationTermsRelativePath, "", "", "", nil, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}
	elems := dom.Find(formatSelectorForAttribute(EllucianDataNameKey, EllucianDataNameValueTerms))
	children := getSingularElement(elems).Children()
	children.Each(func(_ int, s *goquery.Selection) {
		term := sanitizeTerm(s.Text())
		season := parseNameFromTerm(term)
		year := parseYearFromTerm(term)
		ID, _ := s.Attr("value")
		if !containsAnyFromSlice(term, InvalidTerms) {
			res = append(res, entity.Term{
				Season:   season,
				Year:     year,
				TermCode: ID,
			})
		}
	})
	return res
}

// GetSubjects returns a list of all departments / subjects offered for a specific term
func (e *EllucianAPIClient) GetSubjects(request request.SubjectsRequestModel) []entity.Subject {
	res := make([]entity.Subject, 0)
	callingData := map[string]string{
		"p_calling_proc": "bwckschd.p_disp_dyn_sched",
		"p_term":         request.Term,
	}
	// dom := getDocumentModel()
	dom, err := getDocumentModel(request.College, EllucianRegistrationSubjectsRelativePath, EllucianRegistrationTermsRelativePath, "", "", "", callingData, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}
	elems := dom.Find(formatSelectorForAttribute(EllucianDataNameKey, EllucianDataNameValueSubjects))
	children := getSingularElement(elems).Children()
	children.Each(func(_ int, s *goquery.Selection) {
		abbrev, _ := s.Attr("value")
		res = append(res, entity.Subject{
			Abbreviation: abbrev,
			CompleteName: s.Text(),
		})
	})
	return res
}

// GetCourses returns a list of all courses within the specified department/subject
func (e *EllucianAPIClient) GetCourses(request request.CoursesRequestModel) []entity.Course {
	res := make([]entity.Course, 0)
	dom, err := getDocumentModel(request.College, EllucianRegistrationCoursesRelativePath, EllucianRegistrationSubjectsRelativePath, request.Term, request.Subject, "", nil, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}
	elems := dom.Find(formatSelectorForAttribute(EllucianDataClassKey, EllucianDataClassValueCourses))
	elems.Each(func(_ int, s *goquery.Selection) {
		course := parseCourseFromCourseInfo(s.Text())
		res = append(res, course)
	})
	return res
}

// GetSections returns a list of all sections and meeting times for a specific course
func (e *EllucianAPIClient) GetSections(request request.SectionsRequestModel) []entity.Section {
	res := make([]entity.Section, 0)
	dom, err := getDocumentModel(request.College, EllucianRegistrationCoursesRelativePath, EllucianRegistrationSubjectsRelativePath, request.Term, request.Subject, request.Number, nil, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}
	courses := make([]entity.Course, 0)
	elemsCourses := dom.Find(formatSelectorForAttribute(EllucianDataClassKey, EllucianDataClassValueCourses))
	elemsCourses.Each(func(_ int, s *goquery.Selection) {
		course := parseCourseFromCourseInfo(s.Text())
		courses = append(courses, course)
	})
	attrQueryMeetings := formatSelectorForAttribute(EllucianDataClassTableKey, EllucianDataClassTableValueSections)
	elemsMeetings := dom.Find(attrQueryMeetings)
	currentIndex := 0
	elemsMeetings.Each(func(_ int, s *goquery.Selection) {
		sectionMeetings := make([]entity.SectionMeeting, 0)
		tableRows := s.Find(EllucianDataClassTableRowTag)
		tableRowCount := 0
		tableRows.Each(func(_ int, row *goquery.Selection) {
			if tableRowCount > 0 {
				sectionMeetings = append(sectionMeetings, parseScheduledMeetingFromTableRow(row))
			}
			tableRowCount++
		})
		sectionToAdd := entity.Section{
			Course:       courses[currentIndex],
			MeetingTimes: sectionMeetings,
		}
		res = append(res, sectionToAdd)
		currentIndex++
	})
	return res
}

func getDocumentModel(collegeName, relativePath, referrerPath, term, subject, courseNumber string, entries map[string]string, collector *colly.Collector) (*goquery.Document, error) {
	basePage := EllucianUniversitiesDataPages[collegeName]
	dataURL := basePage + relativePath

	var err error
	var dom *goquery.Document
	collector.OnResponse(func(r *colly.Response) {
		dom, err = goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Errorf("Error creating dom from html: %s", err.Error())
		}
	})
	collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Referer", basePage+referrerPath) // CORS spoof
	})

	if len(entries) > 0 || len(term) == 0 {
		err = collector.Post(dataURL, entries)
	} else {
		unencodedData := EllucianCourseDataFormDataDefault + term + EllucianDataFormSubject + subject
		// Add course number to payload if desired
		if len(courseNumber) != 0 {
			unencodedData += (EllucianDataFormCourse + courseNumber)
		} else {
			unencodedData += (EllucianDataFormCourse + "")
		}
		payload := []byte(unencodedData)
		err = collector.PostRaw(dataURL, payload)
	}
	if err != nil {
		log.Errorf("Error posting data to request: %s", err.Error())
		return nil, err
	}

	return dom, nil
}

func defaultCollector() *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Errorf("Error while using colly: %s", err.Error())
	})
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   60 * time.Second,
			KeepAlive: 60 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
	return c
}

func formatSelectorForAttribute(key, value string) string {
	return "*[" + key + "=" + "\"" + value + "\"" + "]"

}

func getSingularElement(elems *goquery.Selection) *goquery.Selection {
	res := elems.First()
	elems.Each(func(_ int, elem *goquery.Selection) {
		val, _ := elem.Attr("value")
		if !strings.EqualFold(val, EllucianDataDummyNode) {
			res = elem
			return
		}
	})
	return res
}

func sanitizeTerm(term string) string {
	return strings.Replace(term, "(View only)", "", -1)
}

func parseNameFromTerm(term string) string {
	if caseInsensitiveContains(term, "Fall") {
		return "Fall"
	} else if caseInsensitiveContains(term, "Winter") {
		return "Winter"
	} else if caseInsensitiveContains(term, "Spring") {
		return "Spring"
	} else if caseInsensitiveContains(term, "Summer") {
		return "Summer"
	}
	return ""
}

func parseYearFromTerm(term string) string {
	return strings.TrimSpace(regexp.MustCompile("[^0-9\\s]").ReplaceAllString(term, ""))
}

func parseScheduledMeetingFromTableRow(row *goquery.Selection) entity.SectionMeeting {
	data := make([]string, 7)
	dataColumns := row.Find(EllucianDataClassTableDataColumnTag)
	dataColumns.Each(func(i int, col *goquery.Selection) {
		data[i] = col.Text()
	})

	return entity.SectionMeeting{
		Type:         data[0],
		Time:         data[1],
		Days:         data[2],
		Location:     data[3],
		DateRange:    data[4],
		ScheduleType: data[5],
		Instructors:  data[6],
	}
}

func parseCourseFromCourseInfo(courseInfo string) entity.Course {
	if strings.Contains(courseInfo, " - ") {
		info := strings.Split(courseInfo, " - ")
		if len(info) != 4 {
			log.Errorf("Error formatting: %s", courseInfo)
		}
		return entity.Course{
			CourseName:    info[2],
			CourseTitle:   info[0],
			Department:    strings.Split(info[2], " ")[0],
			CourseNumber:  strings.Split(info[2], " ")[1],
			CourseID:      info[1],
			CourseSection: info[3],
		}
	}
	log.Errorf("Error formatting: %s", courseInfo)
	return entity.Course{}
}

func caseInsensitiveContains(s, substr string) bool {
	s, substr = strings.ToUpper(s), strings.ToUpper(substr)
	return strings.Contains(s, substr)
}

func containsAnyFromSlice(e string, s []string) bool {
	for _, a := range s {
		if caseInsensitiveContains(a, e) {
			return true
		}
	}
	return false
}
