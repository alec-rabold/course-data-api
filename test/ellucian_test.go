package test

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/alec-rabold/course-data-api/pkg/model/entity"
	"github.com/alec-rabold/course-data-api/pkg/model/request"
	"github.com/alec-rabold/course-data-api/pkg/util/ellucian"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ellucian client", func() {
	It("should successfully implement GetColleges()", func() {
		By("executing GetColleges()")
		Expect(ellucianClient.GetColleges(ctx)).ToNot(BeEmpty())
	})

	// NOTE: as this is not meant to be an exhaustive suite,
	// the tested term, subject, and course are chosen at random
	Context("for each supported college", func() {
		DescribeTable("implements API operations", func(college ellucian.College) {
			rand.Seed(time.Now().UnixNano())

			By("executing GetTerms()")
			terms, err := ellucianClient.GetTerms(ctx, request.Terms{
				College: string(college),
			})
			Expect(err).To(BeNil())
			Expect(terms).ToNot(BeEmpty())

			By("executing GetSubjects()")
			term := getRandomTerm(terms)
			ginkgoLog("\t Term: %s", term.String())
			subjects, err := ellucianClient.GetSubjects(ctx, request.Subjects{
				College: string(college),
				Term:    term.TermCode,
			})
			Expect(err).To(BeNil())
			Expect(subjects).ToNot(BeEmpty())

			By("executing GetCourses()")
			subject := subjects[rand.Intn(len(subjects))]
			ginkgoLog("\t Term: %s ", term.String())
			ginkgoLog("\t Subject: %s", subject.String())
			courses, err := ellucianClient.GetCourses(ctx, request.Courses{
				College: string(college),
				Term:    term.TermCode,
				Subject: subject.Abbreviation,
			})
			Expect(err).To(BeNil())
			Expect(courses).ToNot(BeEmpty())

			By("executing GetSections()")
			course := courses[rand.Intn(len(courses))]
			ginkgoLog("\t Term: %s", term.String())
			ginkgoLog("\t Subject: %s", subject.String())
			ginkgoLog("\t Course: %s", course.String())
			sections, err := ellucianClient.GetSections(ctx, request.Sections{
				College: string(college),
				Term:    term.TermCode,
				Subject: subject.Abbreviation,
				Number:  course.CourseNumber,
			})
			Expect(err).To(BeNil())
			Expect(sections).ToNot(BeEmpty())
		},
			entries()...,
		)
	})
})

func entries() []TableEntry {
	var entries []TableEntry
	for college, name := range ellucian.SupportedColleges {
		entries = append(entries, Entry(name, college))
	}
	return entries
}

// Gets a random term with a bias for the past +/- 1 years
func getRandomTerm(terms []entity.Term) entity.Term {
	var latestTerms []entity.Term
	for _, term := range terms {
		// Ignore if the year isn't as easily parsed (e.g. "Winter 2020-19" then skip it)
		year, err := strconv.Atoi(term.Year)
		if err != nil {
			continue
		}
		if year >= time.Now().Year()-1 && year <= time.Now().Year()+1 {
			latestTerms = append(latestTerms, term)
		}
	}
	// If unable to find any of the latest terms, then
	// just choose one at random from the list of all terms
	if len(latestTerms) == 0 {
		return terms[rand.Intn(len(terms))]
	}
	return latestTerms[rand.Intn(len(latestTerms))]
}

func ginkgoLog(msg string, args ...interface{}) {
	GinkgoWriter.Write([]byte(fmt.Sprintf(msg+"\n", args...)))
}
