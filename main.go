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
	api["leaderboard"] = "https://api.boot.dev/v1/leaderboard_archmage"
	return api[URL]
}
