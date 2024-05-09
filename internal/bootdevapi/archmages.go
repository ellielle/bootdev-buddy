package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/ellielle/bootdev-stats/internal/cache"
)

func GetArchmages(leaderboard string, c cache.Cache) error {

	// Create a new request to https://api.boot.dev/v1/leaderboard_archmage
	req, err := http.NewRequest("GET", leaderboard, nil)
	if err != nil {
		return errors.New(err.Error())
	}

	// Declare `archmages` ahead of time. `archmages`
	// needs to be declared for the JSON from the cache
	// to be unmarshaled
	var archmages = []Archmage{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(leaderboard)
	if ok {
		json.Unmarshal(cacheHit, &archmages)
		// TODO: wait for TUI
		return nil
	}

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New(err.Error())
	}

	// If the request succeeds, we'll marshal
	// the JSON response into `archmages`
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&archmages)
	if err != nil {
		return errors.New(err.Error())
	}

	// create a physical copy of the cache
	// NOTE: this is for debugging and will
	// probably not be kept
	file, err := os.Create("./archmages.json")
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()

	mages, err := json.Marshal(archmages)
	if err != nil {
		return errors.New(err.Error())
	}

	file.Write(mages)

	file.Sync()
	return nil
}
