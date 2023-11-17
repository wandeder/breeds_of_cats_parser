package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/wandeder/breeds_of_cats_parser/app/parser"
	"github.com/wandeder/breeds_of_cats_parser/app/sorter"
	"github.com/wandeder/breeds_of_cats_parser/app/writer"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	link := os.Getenv("LINK")
	if link == "" {
		log.Println("Link is empty")
	}

	//Recursively downloads data about breeds
	unsortedBreeds, err := parser.GetBreedsList(link)
	if err != nil {
		fmt.Println("Error fetching breeds:", err)
	}

	// Sorts breeds by the length of their names in ascending order
	sortedBreeds, err := sorter.GetSortedBreeds(unsortedBreeds)
	if err != nil {
		fmt.Println("Error sorting breeds:", err)
	}

	//Writes data to a JSON file
	writer.WriteBreedsIntoJson(sortedBreeds)
}
