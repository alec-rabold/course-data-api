package service

import (
	"github.com/alec-rabold/EllucianBannerApi-go/model/entity"
	"github.com/alec-rabold/EllucianBannerApi-go/model/request"
	. "github.com/alec-rabold/EllucianBannerApi-go/util"
	// "github.com/anaskhan96/soup"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

// EllucianAPIService is the implementation of the API for universities that use Ellucian Banner
type EllucianAPIService struct {
	collector *colly.Collector
}

// NewEllucianAPIService creates a new instantiation of the API Service for Ellucian universities
func NewEllucianAPIService() *EllucianAPIService {
	return &EllucianAPIService{
		collector: defaultCollector(),
	}
}

// GetTerms returns a list of all academic terms and respective term codes
func (e *EllucianAPIService) GetTerms(request request.TermsRequestModel) {

}

// GetSubjects returns a list of all departments / subjects offered for a specific term
func (e *EllucianAPIService) GetSubjects(request request.SubjectsRequestModel) {

}

// GetCourses returns a list of all courses within the specified department/subject
func (e *EllucianAPIService) GetCourses(request request.CoursesRequestModel) {

}

// GetSections returns a list of all sections and meeting times for a specific course
func (e *EllucianAPIService) GetSections(request request.SectionsRequestModel) []entity.Section {
	res := make([]entity.Section, 0)
	dom, err := getDocumentModel(request.College, EllucianRegistrationCoursesRelativePath, EllucianRegistrationSubjectsRelativePath, request.Term, request.Subject, request.Number, nil, e.collector)
	if err != nil {
		log.Errorf("Error getting document model: %s", err.Error())
	}
	courses := make([]entity.Course, 0)
	attrQueryCourses := formatSelectorForAttribute(EllucianDataClassKey, EllucianDataClassValueCourses)
	elemsCourses := dom.Find(attrQueryCourses)
	elemsCourses.Each(func(i int, s *goquery.Selection) {
		courseInfo := s.Text()
		courses = append(courses, parseCourseFromCourseInfo(courseInfo))
	})
	attrQueryMeetings := formatSelectorForAttribute(EllucianDataClassTableKey, EllucianDataClassTableValueSections)
	// log.Infof("attrQuery: %s", attrQueryMeetings)
	elemsMeetings := dom.Find(attrQueryMeetings)
	// log.Info("HERE")
	currentIndex := 0
	elemsMeetings.Each(func(i int, s *goquery.Selection) {
		sectionMeetings := make([]entity.SectionMeeting, 0)
		// html, err := s.Html()
		tableRows := s.Find(EllucianDataClassTableRowTag)
		tableRowCount := 0
		tableRows.Each(func(i int, row *goquery.Selection) {
			if tableRowCount > 0 {
				sectionMeetings = append(sectionMeetings, parseScheduledMeetingFromTableRow(row))
			}
			tableRowCount++
		})

		currentCourse := courses[currentIndex]
		sectionToAdd := entity.Section{
			CourseID:      currentCourse.CourseID,
			CourseSection: currentCourse.CourseSection,
			MeetingTimes:  sectionMeetings,
		}
		res = append(res, sectionToAdd)
		currentIndex++
	})
	return res
}

func getDocumentModel(collegeName, relativePath, referrerPath, term, subject, courseNumber string, entries []string, collector *colly.Collector) (*goquery.Document, error) {
	log.Debug("GetDocumentMode()")
	basePage := EllucianUniversitiesDataPages[collegeName]
	dataURL := basePage + relativePath
	// e.collector.OnRequest(func(r *colly.Request) {
	// 	r.Headers.Set("Referer", basePage+referrerPath)
	// })
	log.Debug("URL: " + dataURL)
	unencodedData := EllucianCourseDataFormDataDefault + term + EllucianDataFormSubject + subject
	// Add course number to payload if desired
	if len(courseNumber) != 0 {
		unencodedData += (EllucianDataFormCourse + courseNumber)
	} else {
		unencodedData += (EllucianDataFormCourse + "")
	}
	log.Debug("payload: " + unencodedData)
	payload := []byte(unencodedData)

	var err error
	var dom *goquery.Document
	collector.OnResponse(func(r *colly.Response) {
		dom, err = goquery.NewDocumentFromReader(strings.NewReader(string(r.Body)))
		if err != nil {
			log.Errorf("Error creating dom from html: %s", err.Error())
		}
	})

	err = collector.PostRaw(dataURL, payload)
	if err != nil {
		log.Errorf("Error posting data to request: %s", err.Error())
		return nil, err
	}

	return dom, nil
}

func defaultCollector() *colly.Collector {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		log.Debug("OnRequest() called")
		r.Headers.Set("Content-Type", "application/x-www-form-urlencoded")
		// r.Headers.Set("Referer", basePage+referrerPath)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Debug("OnResponse() called")
		log.Debug("reponse: \n============\n", string(r.Body))
		f, err := os.Create("output.html")
		if err != nil {
			log.Errorf("Error opening file: %s", err.Error())
		}
		f.Write(r.Body)
		f.Close()
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Errorf("Error while using colly: %s", err.Error())
	})
	return c
}

func formatSelectorForAttribute(key, value string) string {
	return "*[" + key + "=" + "\"" + value + "\"" + "]"

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
