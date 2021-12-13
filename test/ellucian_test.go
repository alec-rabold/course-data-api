package test

import (
	"math/rand"

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
			By("executing GetTerms()")
			terms, err := ellucianClient.GetTerms(ctx, request.Terms{
				College: string(college),
			})
			Expect(err).To(BeNil())
			Expect(terms).ToNot(BeEmpty())

			By("executing GetSubjects()")
			term := terms[rand.Intn(len(terms))]
			GinkgoWriter.Write([]byte(term.TermCode))
			subjects, err := ellucianClient.GetSubjects(ctx, request.Subjects{
				College: string(college),
				Term:    term.TermCode,
			})
			Expect(err).To(BeNil())
			Expect(subjects).ToNot(BeEmpty())

			By("executing GetCourses()")
			subject := subjects[rand.Intn(len(subjects))]
			GinkgoWriter.Write([]byte(subject.Abbreviation))
			courses, err := ellucianClient.GetCourses(ctx, request.Courses{
				College: string(college),
				Term:    term.TermCode,
				Subject: subject.Abbreviation,
			})
			Expect(err).To(BeNil())
			Expect(courses).ToNot(BeEmpty())

			By("executing GetSections()")
			course := courses[rand.Intn(len(courses))]
			GinkgoWriter.Write([]byte(course.CourseNumber))
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
