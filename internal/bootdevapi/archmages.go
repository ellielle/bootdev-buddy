package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

// Archmages returns a slice of all current Archmages on
// Boot.Dev. It first tries to find a list in the cache, and
// if it doesn't exist or has expired, sends a request for it
func Archmages(c cache.Cache) ([]Archmage, error) {
	// Archmage leaderboard URL
	archmageLB, err := BootDevAPIMap("archmage")
	if err != nil {
		return []Archmage{}, errors.New("error getting archmage url")
	}

	// Declare `archmages` ahead of time. `archmages`
	// needs to be declared for the JSON from the cache
	// to be unmarshaled
	var archmages = []Archmage{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(archmageLB)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &archmages)
		if err != nil {
			return []Archmage{}, err
		}
		// return cache hit and exit early
		return archmages, nil
	}

	// Create a new request
	req, err := http.NewRequest("GET", archmageLB, nil)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
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

	// Marshal the data into JSON to be written to cache
	mages, err := json.Marshal(archmages)
	if err != nil {
		return []Archmage{}, errors.New(err.Error())
	}

	// Write the url and data to the cache so it can be
	// checked before subsuquent requests within the
	// interval period
	c.Add(archmageLB, &mages)

	return archmages, nil
}
