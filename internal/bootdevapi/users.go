package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

func UserInfo(c cache.Cache) (UserData, error) {
	// User information URL
	userURL, err := BootDevAPIMap("user")
	if err != nil {
		return UserData{}, errors.New("error getting user url")
	}

	var user = UserData{}

	// Check cache for a hit before requesting
	cacheHit, ok := c.Get(userURL)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &user)
		if err != nil {
			return UserData{}, err
		}
		// return cache hit and exit early
		return user, nil
	}

	// Create a new request
	req, err := http.NewRequest("POST", userURL, nil)
	if err != nil {
		return UserData{}, errors.New("error creating user request")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return UserData{}, err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&user)
	if err != nil {
		return UserData{}, err
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return UserData{}, err
	}

	// Write user data to the cache
	c.Add(userURL, &userData)

	return user, nil
}
