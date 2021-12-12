package entity

// Term is the data type for a school's specific academic term
type Term struct {
	Season   string `json:"term_season"`
	Year     string `json:"term_year"`
	TermCode string `json:"term_code"`
}
