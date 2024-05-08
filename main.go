package main

import (
	"errors"
	"log"
	"os"
)

type Cache struct {
}

func main() {
	// get the page to crawl from site variable "SITE"
	site := os.Getenv("SITE")

	// if the ENV variable "SITE" doesn't have a valid option,
	// terminate the program
	URL, err := bootDevAPI(site)
	if err != nil {
		log.Fatalf("invalid api option\nvalid options: archmage | stats | daily | karma\n")
	}

	getDBArchmages(URL)
}

// bootDevAPI takes any of the strings below
// as map keys, and returns the url for that endpoint
// archmage = Archmage Leaderboard data
// stats = Global daily stats
// daily = Top daily learners - based on XP earned
// karma: Discord community leaderboard - based on a bunch of Discordiness
func bootDevAPI(URL string) (string, error) {
	api := make(map[string]string)
	api["archmage"] = "https://api.boot.dev/v1/leaderboard_archmage"
	api["stats"] = "https://api.boot.dev/v1/leaderboard_stats"
	api["daily"] = "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"
	api["karma"] = "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"
	if _, ok := api[URL]; !ok {
		return "", errors.New("invalid api")

	}
	return api[URL], nil
}
