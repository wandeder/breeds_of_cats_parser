package app

type (
	Breed struct {
		Breed   string `json:"breed"`
		Country string `json:"country"`
	}
	GetBreedsResponse struct {
		Data []Breed `json:"data"`
	}
	BreedsByCountry struct {
		Country string   `json:"country"`
		Breeds  []string `json:"breeds"`
	}
	Breeds        []Breed
	CountryBreeds map[string][]string
	BreedsList    []BreedsByCountry
)
