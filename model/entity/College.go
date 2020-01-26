package entity

// College is the (meta)data type for individual courses
type College struct {
	NameAbbr string `json:"abbr_name"`
	NameFull string `json:"full_name"`
}
