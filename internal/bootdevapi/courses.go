package bootdevapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ellielle/bootdev-buddy/internal/cache"
)

func CourseList(c cache.Cache, token string) ([]Course, error) {
	coursesURL, err := BootDevAPIMap("courses")
	if err != nil {
		return nil, errors.New("error getting courses url")
	}

	var courses = []Course{}

	// check cache for a hit
	cacheHit, ok := c.Get(coursesURL)
	if ok {
		err := json.Unmarshal([]byte(cacheHit), &courses)
		if err != nil {
			return nil, err
		}
		return courses, nil
	}

	// create new http request
	req, err := http.NewRequest("GET", coursesURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// decode json response into courses if req succeeded
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&courses)
	if err != nil {
		return nil, err
	}

	// Marshal the data into JSON to be written to cache
	cs, err := json.Marshal(courses)
	if err != nil {
		return nil, err
	}

	c.Add(coursesURL, &cs)

	return courses, nil
}
