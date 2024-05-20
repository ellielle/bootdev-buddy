package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

func GetDiscordLeaderboard(c cache.Cache) ([]Archsage, error) {
	//TODO:
	// Community leaderboard URL
	const communityLB = "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"

	// Create a new request
	req, err := http.NewRequest("GET", communityLB, nil)
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}

	// Declare `archsages` ahead of time. `archsages`
	// needs to be declared for the JSON from the cache
	// to be unmarshaled
	var archsages = []Archsage{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(communityLB)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &archsages)
		if err != nil {
			return []Archsage{}, err
		}
		// return cache hit and exit early
		return archsages, nil
	}

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}
	defer resp.Body.Close()

	// If the request succeeds, we'll decode
	// the JSON response into `archsages`
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&archsages)
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}

	// NOTE: this is for debugging and will
	// probably not be kept
	// create a physical copy of the cache
	file, err := os.Create("./archsages.json")
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}
	defer file.Close()

	sages, err := json.Marshal(archsages)
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}

	// Write the url and data to the cache so it can be
	// checked before subsuquent requests within the
	// interval period
	c.Add(communityLB, &sages)

	_, err = file.Write(sages)
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}

	err = file.Sync()
	if err != nil {
		return []Archsage{}, errors.New(err.Error())
	}
	return archsages, nil
}
