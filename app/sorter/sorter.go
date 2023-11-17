package sorter

import (
	"github.com/wandeder/breeds_of_cats_parser/app/models"
	"sort"
)

func sortBreeds(allBreeds models.CountryBreeds) models.BreedsList {
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
	return breedsList

}

func GetSortedBreeds(breeds models.CountryBreeds) (models.BreedsList, error) {
	return sortBreeds(breeds), nil
}
