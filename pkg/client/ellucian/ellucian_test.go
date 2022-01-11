package ellucian

import (
	"testing"

	"github.com/alec-rabold/course-data-api/pkg/model/entity"
	"github.com/stretchr/testify/assert"
)

func TestParseCourseFromCourseInfo(t *testing.T) {
	for name, tt := range map[string]struct {
		in     string
		assert func(*testing.T, entity.Course, error)
	}{
		"standard course format": {
			in: "Introduction to Business - 01234 - BUS 100 - 001",
			assert: func(t *testing.T, course entity.Course, err error) {
				assert.NoError(t, err)
				assert.Equal(t, entity.Course{
					CourseName:    "BUS 100",
					CourseTitle:   "Introduction to Business",
					Department:    "BUS",
					CourseNumber:  "100",
					CourseID:      "01234",
					CourseSection: "001",
				}, course)
			},
		},
		"unexpected course format (too few separators)": {
			in: "Introduction to Agriculture - 56789 - 002",
			assert: func(t *testing.T, course entity.Course, err error) {
				assert.Error(t, err)
			},
		},
		"unexpected course format (too many separators)": {
			in: "Varsity Basketball - Men - 01234 - PHED 113 - 004",
			assert: func(t *testing.T, course entity.Course, err error) {
				assert.NoError(t, err)
				assert.Equal(t, entity.Course{
					CourseName:    "PHED 113",
					CourseTitle:   "Varsity Basketball - Men",
					Department:    "PHED",
					CourseNumber:  "113",
					CourseID:      "01234",
					CourseSection: "004",
				}, course)
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			c, err := parseCourseFromCourseInfo(tt.in)
			tt.assert(t, c, err)
		})
	}
}
