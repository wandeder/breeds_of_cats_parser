package main

import (
	"fmt"
	"github.com/wandeder/breeds_of_cats_parser/app"
	"log"
	"os"

	"github.com/joho/godotenv"
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
	unsortedBreeds, err := app.GetBreedsList(link)
	if err != nil {
		fmt.Println("Error fetching breeds:", err)
	}

	// Sorts breeds by the length of their names in ascending order
	sortedBreeds, err := app.GetSortedBreeds(unsortedBreeds)
	if err != nil {
		fmt.Println("Error sorting breeds:", err)
	}

	//Writes data to a JSON file
	app.WriteBreedsIntoJson(sortedBreeds)
}
