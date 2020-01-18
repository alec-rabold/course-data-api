package request

// SectionsRequestModel contains the required parameters for /sections
type SectionsRequestModel struct {
	College string
	Term    string
	Subject string
	Number  string
}
