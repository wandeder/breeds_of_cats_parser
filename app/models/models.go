package models

type (
	Breed struct {
		Breed   string `json:"breed"`
		Country string `json:"country"`
		Origin  string `json:"origin"`
		Coat    string `json:"coat"`
		Pattern string `json:"pattern"`
	}
	GetBreedsResponse struct {
		CurrentPage int     `json:"current_page"`
		Data        []Breed `json:"data"`
		NextPageUrl string  `json:"next_page_url"`
		Path        string  `json:"path"`
		PerPage     int     `json:"per_page"`
		PrevPageUrl string  `json:"prev_page_url"`
		To          int     `json:"to"`
		Total       int     `json:"total"`
	}
	Breeds []Breed
)
