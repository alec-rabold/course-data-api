package entity

// Section is the data type for individual sections
type Section struct {
	Course
	MeetingTimes []SectionMeeting `json:"meeting_times"`
}

func (s Section) String() string {
	return s.CourseSection
}

// SectionMeeting is the data type for section meeting times
type SectionMeeting struct {
	Type         string `json:"type"`
	Time         string `json:"time"`
	Days         string `json:"days"`
	Location     string `json:"location"`
	DateRange    string `json:"date_range"`
	ScheduleType string `json:"schedule_type"`
	Instructors  string `json:"instructors"`
}
