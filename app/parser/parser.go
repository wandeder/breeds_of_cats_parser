package parser

import (
	"encoding/json"
	"github.com/wandeder/breeds_of_cats_parser/app/models"
	"io"
	"net/http"
)

func fetchBreeds(link string, page int) (models.Breeds, error) {
	response, err := http.Get(link)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res models.GetBreedsResponse
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	breeds := res.Data

	return breeds, nil
}

func parseBreedsRecursive(link string, page int, allBreeds *models.Breeds) error {
	breeds, err := fetchBreeds(link, page)
	if err != nil {
		return err
	}
	// Stop recursion if breeds is empty
	if len(breeds) == 0 {
		return nil
	}

	*allBreeds = append(*allBreeds, breeds...)
	return parseBreedsRecursive(link, page+1, allBreeds)
}

func GetBreedsList(link string) (models.Breeds, error) {
	allBreeds := models.Breeds{}
	if err := parseBreedsRecursive(link, 1, &allBreeds); err != nil {
		return nil, err
	}
	return allBreeds, nil
}
