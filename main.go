package main

import (
	"log"
	"os"
	"time"

	bd "github.com/ellielle/bootdev-stats/internal/bootdevapi"
	"github.com/ellielle/bootdev-stats/internal/cache"
)

type Client struct {
	cache cache.Cache
}

func main() {
	// get the page to crawl from site variable "SITE"
	site := os.Getenv("SITE")
	c := Client{
		cache: cache.NewCache(60 * time.Second),
	}

	// if the ENV variable "SITE" doesn't have a valid option,
	// terminate the program
	URL, err := bd.BootDevAPIMap(site)
	if err != nil {
		log.Fatalf("invalid api option\nvalid options: archmage | stats | daily | karma\n")
	}

	err = bd.GetArchmages(URL, c.cache)
	if err != nil {
		log.Fatalf(err.Error())
	}

}
