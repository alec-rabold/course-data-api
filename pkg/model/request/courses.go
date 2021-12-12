package request

// Courses contains the required input parameters for the "/courses" API handler
type Courses struct {
	College string
	Term    string
	Subject string
}
