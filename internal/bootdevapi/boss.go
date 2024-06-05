package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

// BossBattleStats retrieves stats from the current / most recent boss
// battle happening. Since this is rapidly updating in real time, the
// cache isn't checked before requesting new data.
// It does save to the cache, however, in case it's needed
func BossBattleStats(c cache.Cache, token string) (BossBattle, error) {
	// boss battle leaderboard URL
	bossBattleURL, err := BootDevAPIMap("boss")
	if err != nil {
		return BossBattle{}, errors.New("error getting boss battle url")
	}

	// Create a new request to the boss battle api
	req, err := http.NewRequest("GET", bossBattleURL, nil)
	if err != nil {
		return BossBattle{}, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	// Send request for boss battle info
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return BossBattle{}, err
	}
	defer resp.Body.Close()

	// If request succeeds, decode the JSON
	// response into `boss`
	var boss BossBattle

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&boss)
	if err != nil {
		return BossBattle{}, err
	}

	bossData, err := json.Marshal(resp.Body)
	if err != nil {
		return BossBattle{}, err
	}

	// add to cache just in case the data is needed sooner than it is
	// fetched again.
	c.Add(bossBattleURL, &bossData)

	return boss, nil
}
