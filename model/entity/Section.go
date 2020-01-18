package entity

// Section is the (meta)data type for individual sections
type Section struct {
	CourseID      string   `json:"course_id"`
	CourseSection string   `json:"section"`
	MeetingTimes  []string `json:"meeting_times"`
}
