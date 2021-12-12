package entity

// Subject is the data type for a specific course subject
type Subject struct {
	Abbreviation string `json:"subj_abbr"`
	CompleteName string `json:"subj_name"`
}
