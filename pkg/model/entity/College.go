package entity

import (
	"fmt"
)

// College is the data type supported colleges
type College struct {
	NameAbbr string `json:"abbr_name"`
	NameFull string `json:"full_name"`
}

func (c College) String() string {
	return fmt.Sprintf("%s - %s", c.NameAbbr, c.NameFull)
}
