package entity

// College is the (meta)data type supported colleges
type College struct {
	NameAbbr string `json:"abbr_name"`
	NameFull string `json:"full_name"`
}
