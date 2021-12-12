package request

// Sections contains the required parameters for the "/sections" API handler
type Sections struct {
	College string
	Term    string
	Subject string
	Number  string
}
