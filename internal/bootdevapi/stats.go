package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

func GetGeneralStats(c cache.Cache) {
	// Regular leaderboard stats URL
	const statsLB = "https://api.boot.dev/v1/leaderboard_stats"

	// Create a new request to https://api.boot.dev/v1/leaderboard_archmage
	req, err := http.NewRequest("GET", statsLB, nil)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}

	// Declare `archmages` ahead of time. `archmages`
	// needs to be declared for the JSON from the cache
	// to be unmarshaled
	var archmages = []Archmage{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(statsLB)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &archmages)
		if err != nil {
			return []Archmage{}, err
		}
		// return cache hit and exit early
		return archmages, nil
	}

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}
	defer resp.Body.Close()

	// If the request succeeds, we'll decode
	// the JSON response into `archmages`
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&archmages)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}

	// NOTE: this is for debugging and will
	// probably not be kept
	// create a physical copy of the cache
	file, err := os.Create("./archmages.json")
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}
	defer file.Close()

	mages, err := json.Marshal(archmages)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}

	// Write the url and data to the cache so it can be
	// checked before subsuquent requests within the
	// interval period
	c.Add(statsLb, &mages)

	_, err = file.Write(mages)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}

	err = file.Sync()
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}
	return archmages, nil
}
