package main

import (
	"os"
)

func main() {
	// get the page to crawl from site variable "SITE"
	site := os.Getenv("SITE")
	getBDArchmages(site)
	// getBDLeaderboard(env)
}
