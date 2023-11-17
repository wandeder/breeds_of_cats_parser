package sorter

import (
	"github.com/wandeder/breeds_of_cats_parser/app/models"
	"sort"
)

func GetSortedBreeds(allBreeds models.CountryBreeds) (models.BreedsList, error) {
	// Sorts breeds by the length of their names in ascending order

	breedsList := models.BreedsList{}
	for country, breeds := range allBreeds {
		if country == "" {
			country = "Undefined"
		}

		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]) < len(breeds[j])
		})

		sortedBreeds := models.BreedsByCountry{
			Country: country,
			Breeds:  breeds,
		}

		breedsList = append(breedsList, sortedBreeds)
	}
	return breedsList, nil
}
