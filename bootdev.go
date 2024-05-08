package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func check(e error, message string) {
	if e != nil {
		log.Fatalf(message, e.Error())
	}
}

func getDBArchmages(leaderboard string) {
	// Create a new request to https://api.boot.dev/v1/leaderboard_archmage
	req, err := http.NewRequest("GET", leaderboard, nil)
	check(err, "Something is wrong with the request URL, %v")

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	check(err, "Something went during the request: %v")

	// If the request succeeds, we'll marshal
	// the JSON response into `archmages`
	archmages := []Archmage{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&archmages)
	check(err, "Error marshaling JSON data: %v")

	log.Println(archmages[0].Handle)

	// FIXME: decoder has value, convert to json to see what you're working with
	file, err := os.Create("archmages.json")
	check(err, "Error creating archmages.json: %v")
	defer file.Close()

}
