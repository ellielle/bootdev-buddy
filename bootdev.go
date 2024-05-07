package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getDBArchmages(leaderboard string) {
	// Create a new request to https://api.boot.dev/v1/leaderboard_archmage
	req, err := http.NewRequest("GET", leaderboard, nil)
	if err != nil {
		log.Fatalf("Something is wrong with the request URL, %v", err.Error())
	}

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Something went during the request: %v", err.Error())
	}

	// If the request succeeds, we'll marshal
	// the JSON response into `archmages`
	archmages := []Archmage{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&archmages)
	if err != nil {
		log.Fatalf("Error marshaling JSON data: %v", err.Error())
	}

	log.Println(archmages[0].Handle)

	// FIXME: decoder has value, convert to json to see what you're working with
}
