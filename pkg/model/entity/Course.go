package entity

import (
	"fmt"
)

// Course is the data type for individual courses
type Course struct {
	CourseName    string `json:"course_name"`
	CourseTitle   string `json:"course_title"`
	Department    string `json:"department"`
	CourseNumber  string `json:"course_number"`
	CourseID      string `json:"course_id"`
	CourseSection string `json:"-"`
}

func (c Course) String() string {
	return fmt.Sprintf("%s - %s (%s)", c.CourseName, c.CourseTitle, c.CourseID)
}
