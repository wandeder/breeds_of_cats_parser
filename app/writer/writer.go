package writer

import (
	"fmt"
	"github.com/wandeder/breeds_of_cats_parser/app/models"
)

func PrintAllBreeds(allBreeds models.Breeds) {
	for _, breed := range allBreeds {
		fmt.Printf(
			"Breed: %s, Country: %s, Origin: %s, Coat: %s, Pattern: %s\n",
			breed.Breed,
			breed.Country,
			breed.Origin,
			breed.Coat,
			breed.Pattern,
		)
	}

}
