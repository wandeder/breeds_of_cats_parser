package app

import (
	"sort"
)

func GetSortedBreeds(allBreeds CountryBreeds) (BreedsList, error) {
	// Sorts breeds by the length of their names in ascending order

	breedsList := BreedsList{}
	for country, breeds := range allBreeds {
		if country == "" {
			country = "Undefined"
		}

		sort.Slice(breeds, func(i, j int) bool {
			return len(breeds[i]) < len(breeds[j])
		})

		sortedBreeds := BreedsByCountry{
			Country: country,
			Breeds:  breeds,
		}

		breedsList = append(breedsList, sortedBreeds)
	}
	return breedsList, nil
}
