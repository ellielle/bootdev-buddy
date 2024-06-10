package bootdevapi

import (
	"encoding/json"
	"errors"
	"log"
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

func CoursesProgress(c cache.Cache, token string) (CourseProgress, error) {
	progressURL, err := BootDevAPIMap("courses_progress")
	if err != nil {
		return CourseProgress{}, err
	}

	var progress CourseProgress

	// check cache for a hit
	cacheHit, ok := c.Get(progressURL)
	if ok {
		log.Print("cache hit on progress")
		err := json.Unmarshal([]byte(cacheHit), &progress)
		if err != nil {
			return CourseProgress{}, err
		}

		return progress, nil
	}

	// create a new http request
	req, err := http.NewRequest("GET", progressURL, nil)
	if err != nil {
		return CourseProgress{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return CourseProgress{}, err
	}
	defer resp.Body.Close()

	// decode json response into progress if req succeeded
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&progress)
	if err != nil {
		return CourseProgress{}, err
	}

	// Marshal data into JSON to be written to cache
	prog, err := json.Marshal(progress)
	if err != nil {
		return CourseProgress{}, err
	}

	c.Add(progressURL, &prog)

	return progress, nil
}
