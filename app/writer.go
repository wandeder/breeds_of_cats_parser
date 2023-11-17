package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func WriteBreedsIntoJson(allBreeds BreedsList) error {
	//Writes data to a JSON file

	fileName := os.Getenv("FILE_NAME")
	if fileName == "" {
		log.Println("File name is empty")
	}

	file, err := json.MarshalIndent(&allBreeds, "", "  ")
	if err != nil {
		fmt.Println("Error Json serializing:", err)
	}

	err = ioutil.WriteFile(fileName, file, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}

	return nil
}
