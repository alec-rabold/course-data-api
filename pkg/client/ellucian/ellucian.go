package ellucian

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/alec-rabold/course-data-api/pkg/crawler"
	"github.com/alec-rabold/course-data-api/pkg/model/entity"
	"github.com/alec-rabold/course-data-api/pkg/model/request"
	"github.com/alec-rabold/course-data-api/pkg/util/ellucian"
	"github.com/hashicorp/go-multierror"
)

// var _ client.Client = &Client{}

// Client is the implementation of the API for universities that use ellucian. Banner
type Client struct {
	crawler *crawler.Crawler
}

// NewClient creates a new instantiation of the API Service for ellucian. universities
func NewClient(c *crawler.Crawler) *Client {
	return &Client{
		crawler: c,
	}
}

// GetColleges returns a list of all supported colleges to be used for this API
func (c *Client) GetColleges(ctx context.Context) ([]entity.College, error) {
	var res []entity.College
	for key, value := range ellucian.SupportedColleges {
		res = append(res, entity.College{
			NameAbbr: string(key),
			NameFull: value,
		})
	}

	return res, nil
}

// GetTerms returns a list of all academic terms and respective term codes
func (c *Client) GetTerms(ctx context.Context, request request.Terms) ([]entity.Term, error) {
	var res []entity.Term

	base := ellucian.SelfServicePages[ellucian.College(request.College)]
	doc, err := c.crawler.GetDocument(ctx, fmt.Sprintf("%s/%s", base, ellucian.RegistrationTermsProcedure))
	if err != nil {
		return nil, fmt.Errorf("error getting document model: %w", err)
	}

	elems := doc.Find(attributeSelector(ellucian.DataNameKey, ellucian.DataNameValueTerms))
	selection, err := getNonDummySelection(elems)
	if err != nil {
		return nil, fmt.Errorf("error getting non-dummy selection: %w", err)
	}
	children := selection.Children()
	if err != nil {
		return nil, fmt.Errorf("no child elements for selection %v", selection)
	}
	children.Each(func(_ int, s *goquery.Selection) {
		term := sanitizeTerm(s.Text())
		season := parseNameFromTerm(term)
		year := parseYearFromTerm(term)
		ID, exists := s.Attr("value")
		if !exists {
			return
		}
		if !containsAnyFromSlice(term, ellucian.InvalidTerms) {
			res = append(res, entity.Term{
				Season:   season,
				Year:     year,
				TermCode: ID,
			})
		}
	})

	return res, nil
}

// GetSubjects returns a list of all departments / subjects offered for a specific term
func (c *Client) GetSubjects(ctx context.Context, request request.Subjects) ([]entity.Subject, error) {
	var res []entity.Subject

	base := ellucian.SelfServicePages[ellucian.College(request.College)]
	doc, err := c.crawler.GetDocument(ctx, fmt.Sprintf("%s/%s", base, ellucian.RegistrationSubjectsProcedure),
		crawler.WithData(url.Values{
			"p_term":         []string{request.Term},
			"p_calling_proc": []string{"bwckschd.p_disp_dyn_sched"},
		}),
		crawler.WithReferer(fmt.Sprintf("%s/%s", base, ellucian.RegistrationTermsProcedure)),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting document model: %w", err)
	}

	elems := doc.Find(attributeSelector(ellucian.DataNameKey, ellucian.DataNameValueSubjects))
	selection, err := getNonDummySelection(elems)
	if err != nil {
		return nil, fmt.Errorf("error getting non-dummy selection: %w", err)
	}
	selection.Children().Each(func(_ int, s *goquery.Selection) {
		abbrev, exists := s.Attr("value")
		if !exists {
			return
		}
		res = append(res, entity.Subject{
			Abbreviation: abbrev,
			CompleteName: s.Text(),
		})
	})

	return res, nil
}

// GetCourses returns a list of all courses within the specified department/subject
func (c *Client) GetCourses(ctx context.Context, request request.Courses) ([]entity.Course, error) {
	var res []entity.Course

	base := ellucian.SelfServicePages[ellucian.College(request.College)]
	data := ellucian.DefaultCourseDataFormValues()
	data.SetTerm(request.Term)
	data.SetSubject(request.Subject)

	doc, err := c.crawler.GetDocument(ctx, fmt.Sprintf("%s/%s", base, ellucian.RegistrationCoursesProcedure),
		crawler.WithData(data.Values()),
		crawler.WithReferer(fmt.Sprintf("%s/%s", base, ellucian.RegistrationTermsProcedure)),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting document model: %w", err)
	}

	elems := doc.Find(attributeSelector(ellucian.DataClassKey, ellucian.DataClassValueCourses))
	var parseErr error
	elems.Each(func(_ int, s *goquery.Selection) {
		c, err := parseCourseFromCourseInfo(s.Text())
		if err != nil {
			parseErr = multierror.Append(parseErr, err)
		}
		res = append(res, c)
	})
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing course info: %w", parseErr)
	}

	return res, nil
}

// GetSections returns a list of all sections and meeting times for a specific course
func (c *Client) GetSections(ctx context.Context, request request.Sections) ([]entity.Section, error) {
	var res []entity.Section

	base := ellucian.SelfServicePages[ellucian.College(request.College)]
	data := ellucian.DefaultCourseDataFormValues()
	// TODO: these constants should live in package "util/ellucian./constants.go"
	// TODO: see if we can use a struct, serialize as json, and set the content encoding as such
	data.SetTerm(request.Term)
	data.SetSubject(request.Subject)
	data.SetCourse(request.Number)

	doc, err := c.crawler.GetDocument(ctx, fmt.Sprintf("%s/%s", base, ellucian.RegistrationCoursesProcedure),
		crawler.WithData(data.Values()),
		crawler.WithReferer(fmt.Sprintf("%s/%s", base, ellucian.RegistrationSubjectsProcedure)),
	)
	if err != nil {
		return nil, fmt.Errorf("error getting document model: %w", err)
	}

	// TODO: maybe just call GetCourses() here instead
	var courses []entity.Course
	elemsCourses := doc.Find(attributeSelector(ellucian.DataClassKey, ellucian.DataClassValueCourses))

	var parseErr error
	elemsCourses.Each(func(_ int, s *goquery.Selection) {
		c, err := parseCourseFromCourseInfo(s.Text())
		if err != nil {
			parseErr = multierror.Append(parseErr, err)
		}
		courses = append(courses, c)
	})
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing course info: %w", parseErr)
	}
	elemsMeetings := doc.Find(attributeSelector(ellucian.DataClassTableKey, ellucian.DataClassTableValueSections))
	elemsMeetings.Each(func(idx int, s *goquery.Selection) {
		var sectionMeetings []entity.SectionMeeting
		tableRows := s.Find(ellucian.DataClassTableRowTag)
		tableRows.Each(func(tableRowIdx int, row *goquery.Selection) {
			// Don't process table headers (headers have the same html structure as data rows)
			if tableRowIdx == 0 {
				return
			}
			m, err := parseScheduledMeetingFromTableRow(row)
			if err != nil {
				parseErr = multierror.Append(parseErr, err)
			}
			sectionMeetings = append(sectionMeetings, m)
		})
		// TODO: this technically works but is pretty suspect in assuming the course index
		sectionToAdd := entity.Section{
			Course:       courses[idx],
			MeetingTimes: sectionMeetings,
		}
		res = append(res, sectionToAdd)
	})
	if parseErr != nil {
		return nil, fmt.Errorf("error parsing course info: %w", parseErr)
	}

	return res, nil
}

// Many ellucian. selections contain dummy nodes so we need to filter those out
func getNonDummySelection(elems *goquery.Selection) (res *goquery.Selection, err error) {
	elems.Each(func(_ int, elem *goquery.Selection) {
		val, _ := elem.Attr("value")
		if strings.EqualFold(val, ellucian.DataDummyNode) {
			return
		}
		// Ensure there's no more than one non-dummy selection,
		// otherwise we'll need to come back and update this logic
		if res != nil {
			err = fmt.Errorf("more than one non-dummy node encountered")
		}
		res = elem
	})
	if res == nil {
		return nil, fmt.Errorf("no non-dummy nodes found")
	}
	return res, err
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

func parseScheduledMeetingFromTableRow(row *goquery.Selection) (entity.SectionMeeting, error) {
	var data []string
	dataColumns := row.Find(ellucian.DataClassTableDataColumnTag)
	dataColumns.Each(func(_ int, col *goquery.Selection) {
		data = append(data, col.Text())
	})
	// TODO: This function shouldn't be so rigid (e.g. sometimes there's an extra <td> at the end for textbook info)
	if len(data) < 7 {
		return entity.SectionMeeting{}, fmt.Errorf("unexpected section meeting data in %v (len %d)", data, len(data))
	}

	return entity.SectionMeeting{
		Type:         data[0],
		Time:         data[1],
		Days:         data[2],
		Location:     data[3],
		DateRange:    data[4],
		ScheduleType: data[5],
		Instructors:  data[6],
	}, nil
}

func parseCourseFromCourseInfo(courseInfo string) (entity.Course, error) {
	info := strings.Split(courseInfo, " - ")
	// The expected length is 4. If the actual length is less, than error out as there's not enough to parse.
	if len(info) < 4 {
		return entity.Course{}, fmt.Errorf("not enough information to parse course: %v (len %d)", info, len(info))
	}
	// If there are more than the expected number of hyphens, then assume it's due to the course name at the beginning
	// (e.g. "Varsity Basketball - Men - 2025 - PHED 1013 - A" where "Varsity Basketball - Men" is the course name)
	if len(info) > 4 {
		info[0] = strings.Join(info[:len(info)-3], " - ")
		for i := 1; i < 4; i++ {
			info[i] = info[len(info)-4+i]
		}
		info = info[:4]
	}
	return entity.Course{
		CourseName:    info[len(info)-2],
		CourseTitle:   strings.Join(info[:len(info)-3], ""),
		Department:    strings.Split(info[len(info)-2], " ")[0],
		CourseNumber:  strings.Split(info[len(info)-2], " ")[1],
		CourseID:      info[len(info)-3],
		CourseSection: info[len(info)-1],
	}, nil
}

func caseInsensitiveContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

func containsAnyFromSlice(e string, s []string) bool {
	for _, a := range s {
		if caseInsensitiveContains(a, e) {
			return true
		}
	}
	return false
}

// Formats a basic goquery selector for the given key/value pair
func attributeSelector(key, value string) string {
	return fmt.Sprintf(`*[%s="%s"]`, key, value)
}
