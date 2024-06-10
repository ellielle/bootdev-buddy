package main

import (
	"log"

	"github.com/ellielle/bootdev-buddy/internal/bootdevapi"
)

func (a *App) Courses() []bootdevapi.Course {
	list, err := bootdevapi.CourseList(a.cache, a.tokens.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	return list
}

func (a *App) CoursesProgress() bootdevapi.CourseProgress {
	list, err := bootdevapi.CoursesProgress(a.cache, a.tokens.AccessToken)
	if err != nil {
		log.Fatal(err)
	}

	return list
}
