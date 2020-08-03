package client

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/entity"
	"github.com/alec-rabold/EllucianBannerApi-go/pkg/model/request"
	. "github.com/alec-rabold/EllucianBannerApi-go/pkg/util"
	"github.com/alec-rabold/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
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
	set := make(map[string]entity.Course)
	dom, err := getDocumentModel(request.College, EllucianRegistrationCoursesRelativePath, EllucianRegistrationSubjectsRelativePath, request.Term, request.Subject, "", nil, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}

	elems := dom.Find(formatSelectorForAttribute(EllucianDataClassKey, EllucianDataClassValueCourses))
	elems.Each(func(_ int, s *goquery.Selection) {
		course := parseCourseFromCourseInfo(s.Text())
		if _, exists := set[course.CourseName]; !exists {
			set[course.CourseName] = course
		}
	})
	res := make([]entity.Course, 0, len(set))
	for _, v := range set {
		res = append(res, v)
	}
	// try different parsing format if nothing found (HACKY)
	if len(res) == 0 {
		elems = dom.Find("tr")
		fmt.Printf("\nnum of tr's found: %d\n", elems.Length())
		res = parseCoursesFromCourseInfoRows(elems, dom)
		for i, course := range res {
			fmt.Printf("Course #%d: \n%v\n\n", i, course)
		}
	}
	// elems.Each(func(_ int, s *goquery.Selection) {
	// 	course := parseCourseFromCourseInfoRows(s.Text())
	// 	fmt.Printf("Course: %v", course)
	// 	if _, exists := set[course.CourseName]; !exists {
	// 		set[course.CourseName] = course
	// 	}
	// })
	return res
}

// GetSections returns a list of all sections and meeting times for a specific course
func (e *EllucianAPIClient) GetSections(req request.SectionsRequestModel) []entity.Section {
	res := make([]entity.Section, 0)
	dom, err := getDocumentModel(req.College, EllucianRegistrationCoursesRelativePath, EllucianRegistrationSubjectsRelativePath, req.Term, req.Subject, req.Number, nil, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}
	courses := e.GetCourses(request.CoursesRequestModel{
		College: req.College,
		Term:    req.Term,
		Subject: req.Subject,
	})
	// courses := make([]entity.Course, 0)
	// elemsCourses := dom.Find(formatSelectorForAttribute(EllucianDataClassKey, EllucianDataClassValueCourses))
	// elemsCourses.Each(func(_ int, s *goquery.Selection) {
	// 	course := parseCourseFromCourseInfo(s.Text())
	// 	courses = append(courses, course)
	// })
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
	// try different parsing format if nothing found (HACKY)
	if len(res) == 0 {
		elemsMeetings = dom.Find("tr")
		// fmt.Printf("\nnum of tr's found: %d\n", elems.Length())
		res = parseSectionsFromCourseInfoRows(elemsMeetings, dom)
		for i, section := range res {
			fmt.Printf("Section #%d: \n%v\n\n", i, section)
		}
	}
	return res
}

func getDocumentModel(collegeName, relativePath, referrerPath, term, subject, courseNumber string, entries map[string]string, collector *colly.Collector) (*goquery.Document, error) {
	basePage := EllucianUniversitiesDataPages[collegeName]
	dataURL := basePage + relativePath

	var err error
	var dom *goquery.Document
	collector.OnResponse(func(r *colly.Response) {
		// fmt.Println(string(r.Body))
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
		// TODO: parse the <input> names/values from the referrer
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
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.246")
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
		ResponseHeaderTimeout: 30 * time.Second,
	})
	c.SetRequestTimeout(30 * time.Second)
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
			log.Debugf("Parsed course info has more than 3 hyphens: %s", courseInfo)
		}
		return entity.Course{
			// Parse info from back to front
			// because course title may include hyphen
			CourseName:    info[len(info)-2],
			CourseTitle:   strings.Join(info[:len(info)-3], ""),
			Department:    strings.Split(info[len(info)-2], " ")[0],
			CourseNumber:  strings.Split(info[len(info)-2], " ")[1],
			CourseID:      info[len(info)-3],
			CourseSection: info[len(info)-1],
		}
	}
	log.Errorf("Error formatting: %s", courseInfo)
	return entity.Course{}
}

// Sooo hacky, but that's the fun of parsing variable-structured data
func parseCoursesFromCourseInfoRows(tableRows *goquery.Selection, document *goquery.Document) []entity.Course {
	set := make(map[string]entity.Course)
	for _, trNode := range tableRows.Nodes {
		var parsedCourse entity.Course
		// Create single selection
		tableRow := goquery.NewSingleSelection(trNode, document)
		tableDataSelection := tableRow.Find("td")
		colNum := 0
		for cur, tdNode := range tableDataSelection.Nodes {
			// Need to skip the top-left <td></td>
			if cur == 0 {
				continue
			}
			tableData := goquery.NewSingleSelection(tdNode, document)
			// Check if <tr> children (<td>) contain CLASS="dddefault"
			if tdAttrClass, exists := tableData.Attr("class"); exists && tdAttrClass == "dddefault" {
				// check if <tr> is for related section; if so, break
				if tableData.Text() == "&nbsp;" {
					fmt.Printf("\nSkipped: &nbsp;\n")
					break
				}
				switch colNum {
				case 0:
					parsedCourse.CourseID = strings.TrimSpace(tableData.Text())
					break
				case 1:
					parsedCourse.Department = strings.TrimSpace(tableData.Text())
					break
				case 2:
					parsedCourse.CourseNumber = strings.TrimSpace(tableData.Text())
					break
				case 3:
					parsedCourse.CourseSection = strings.TrimSpace(tableData.Text())
					break
				case 5:
					parsedCourse.CourseTitle = strings.TrimSpace(tableData.Text())
				}
				colNum++
			} else {
				fmt.Print("\nSkipped, didn't contains dddefault\n")
				break
			}
		}
		parsedCourse.CourseName = fmt.Sprintf("%s %s", parsedCourse.Department, parsedCourse.CourseNumber)
		// Check for zero-value object
		if parsedCourse.IsComplete() {
			if _, exists := set[parsedCourse.CourseName]; !exists {
				set[parsedCourse.CourseName] = parsedCourse
			}
		}
	}
	res := make([]entity.Course, 0, len(set))
	for _, v := range set {
		res = append(res, v)
	}
	return res
}

func parseSectionsFromCourseInfoRows(tableRows *goquery.Selection, document *goquery.Document) []entity.Section {
	set := make(map[string]entity.Section)
	for _, trNode := range tableRows.Nodes {
		var parsedSection entity.Section
		// Create single selection
		tableRow := goquery.NewSingleSelection(trNode, document)
		tableDataSelection := tableRow.Find("td")
		colNum := 0
		for cur, tdNode := range tableDataSelection.Nodes {
			// Need to skip the top-left <td></td>
			if cur == 0 {
				continue
			}
			tableData := goquery.NewSingleSelection(tdNode, document)
			// Check if <tr> children (<td>) contain CLASS="dddefault"
			if tdAttrClass, exists := tableData.Attr("class"); exists && tdAttrClass == "dddefault" {
				// check if <tr> is for related section; if so, break
				if tableData.Text() == "&nbsp;" {
					fmt.Printf("\nSkipped: &nbsp;\n")
					break
				}
				switch colNum {
				case 0:
					parsedSection.CourseID = strings.TrimSpace(tableData.Text())
					break
				case 1:
					parsedSection.Department = strings.TrimSpace(tableData.Text())
					break
				case 2:
					parsedSection.CourseNumber = strings.TrimSpace(tableData.Text())
					break
				case 3:
					parsedSection.CourseSection = strings.TrimSpace(tableData.Text())
					break
				case 5:
					parsedSection.CourseTitle = strings.TrimSpace(tableData.Text())
				}
				colNum++
			} else {
				fmt.Print("\nSkipped, didn't contains dddefault\n")
				break
			}
		}
		parsedSection.CourseName = fmt.Sprintf("%s %s", parsedSection.Department, parsedSection.CourseNumber)
		// Check for zero-value object
		if parsedSection.IsComplete() {
			if _, exists := set[parsedSection.CourseName]; !exists {
				set[parsedSection.CourseName] = parsedSection
			}
		}
	}
	res := make([]entity.Section, 0, len(set))
	for _, v := range set {
		res = append(res, v)
	}
	return res
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
