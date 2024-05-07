package main

import (
	"os"
)

func main() {
	// get the page to crawl from site variable "SITE"
	site := os.Getenv("SITE")

	getDBArchmages(site)
}

func bootDevAPI(URL string) string {
	api := make(map[string]string)
	api["archmage"] = "https://api.boot.dev/v1/leaderboard_archmage"
	api["stats"] = "https://api.boot.dev/v1/leaderboard_stats"
	api["daily"] = "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"
	api["karma"] = "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"
	return api[URL]
}
