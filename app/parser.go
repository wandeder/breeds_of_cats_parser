package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func fetchBreeds(link string, page int) (Breeds, error) {
	url := fmt.Sprintf("%s?page=%d", link, page)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var res GetBreedsResponse
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

func parseBreedsRecursive(link string, page int, allBreeds CountryBreeds) error {
	breeds, err := fetchBreeds(link, page)
	if err != nil {
		return err
	}
	// Stop recursion if breeds is empty
	if len(breeds) == 0 {
		return nil
	}
	for _, breed := range breeds {
		allBreeds[breed.Country] = append(allBreeds[breed.Country], breed.Breed)
	}
	return parseBreedsRecursive(link, page+1, allBreeds)
}

func GetBreedsList(link string) (CountryBreeds, error) {
	//Recursively downloads data about breeds

	allBreeds := CountryBreeds{}
	if err := parseBreedsRecursive(link, 1, allBreeds); err != nil {
		return nil, err
	}
	return allBreeds, nil
}
