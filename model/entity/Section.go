package entity

// Section is the (meta)data type for individual sections
type Section struct {
	CourseID      string           `json:"course_id"`
	CourseSection string           `json:"section"`
	MeetingTimes  []SectionMeeting `json:"meeting_times"`
}
