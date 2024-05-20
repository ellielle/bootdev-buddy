package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

func GetGeneralStats(c cache.Cache) (GlobalStats, error) {
	// Regular leaderboard stats URL
	const statsLB = "https://api.boot.dev/v1/leaderboard_stats"

	// Create a new request to https://api.boot.dev/v1/leaderboard_stats
	req, err := http.NewRequest("GET", statsLB, nil)
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}

	// Declare `globalStats` ahead of time. `globalStats`
	// needs to be declared for the JSON from the cache
	// to be unmarshaled
	var globalStats = GlobalStats{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(statsLB)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &globalStats)
		if err != nil {
			return GlobalStats{}, err
		}
		// return cache hit and exit early
		return globalStats, nil
	}

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}
	defer resp.Body.Close()

	// If the request succeeds, we'll decode
	// the JSON response into `globalStats`
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&globalStats)
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}

	// NOTE: this is for debugging and will
	// probably not be kept
	// create a physical copy of the cache
	file, err := os.Create("./stats.json")
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}
	defer file.Close()

	statsJSON, err := json.Marshal(globalStats)
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}

	// Write the url and data to the cache so it can be
	// checked before subsuquent requests within the
	// interval period
	c.Add(statsLB, &statsJSON)

	_, err = file.Write(statsJSON)
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}

	err = file.Sync()
	if err != nil {
		return GlobalStats{}, errors.New(err.Error())
	}
	return globalStats, nil
}
