package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getDBArchmages(leaderboard string) {
	// create a new request to https://api.boot.dev/v1/leaderboard_archmage
	req, err := http.Get(leaderboard)
	if err != nil {
		log.Fatalf("Something is wrong with the request URL, %v", err.Error())
	}

	archmages := []Archmage{}
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&archmages)
	if err != nil {
		log.Fatalf("Error marshaling JSON data: %v", err.Error())
	}

	log.Println(decoder)

	// FIXME: decoder has value, convert to json to see what you're working with
}
