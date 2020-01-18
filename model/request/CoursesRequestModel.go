package request

// CoursesRequestModel contains the required parameters for /courses
type CoursesRequestModel struct {
	College string
	Term    string
	Subject string
}
