package sorter

import (
	"github.com/wandeder/breeds_of_cats_parser/app/models"
	"sort"
)

func sortedByName(allBreeds models.CountryBreeds) models.BreedsList {
	breedsList := models.BreedsList{}
	for country, breed := range allBreeds {
		sort.Strings(breed)
		sortedBreeds := models.BreedsByCountry{
			Country: country,
			Breeds:  breed,
		}
		breedsList = append(breedsList, sortedBreeds)
	}
	return breedsList

}

func GetSortedBreeds(breeds models.CountryBreeds) (models.BreedsList, error) {
	return sortedByName(breeds), nil
}
