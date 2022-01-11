package entity

import (
	"fmt"
)

// Term is the data type for a school's specific academic term
type Term struct {
	Season   string `json:"term_season"`
	Year     string `json:"term_year"`
	TermCode string `json:"term_code"`
}

func (t Term) String() string {
	return fmt.Sprintf("%s %s (%s)", t.Season, t.Year, t.TermCode)
}
