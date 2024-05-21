package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

// DailyStats returns a slice of the top 30 learners
// from Boot.Dev's leaderboard page.
// As this won't change often, the refresh time on this can be low.
// TODO:
// The list will be paginated to only show 10
func GetDailyStats(c cache.Cache) ([]LeaderboardUser, error) {
	// daily stats leaderboard URL
	const dailyStatsLB = "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"

	// TODO:

	// Create a new request
	req, err := http.NewRequest("GET", dailyStatsLB, nil)
	if err != nil {
		return []LeaderboardUser{}, errors.New(err.Error())
	}

	// Declare `dailyStats` ahead of time. `dailyStats`
	// needs to be declared for the JSON from the cache
	// to be unmarshaled
	var dailyStats = []LeaderboardUser{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(dailyStatsLB)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &dailyStats)
		if err != nil {
			return []LeaderboardUser{}, err
		}
		// return cache hit and exit early
		return dailyStats, nil
	}

	// Send that request out!
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []LeaderboardUser{}, errors.New(err.Error())
	}
	defer resp.Body.Close()

	// If the request succeeds, we'll decode
	// the JSON response into `dailyStats`
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&dailyStats)
	if err != nil {
		return []LeaderboardUser{}, errors.New(err.Error())
	}

	mages, err := json.Marshal(dailyStats)
	if err != nil {
		return []LeaderboardUser{}, errors.New(err.Error())
	}

	// Write the url and data to the cache so it can be
	// checked before subsuquent requests within the
	// interval period
	c.Add(dailyStatsLB, &mages)

	return dailyStats, nil
}
