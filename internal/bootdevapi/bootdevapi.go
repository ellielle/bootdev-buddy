package bootdevapi

import "errors"

// BootDevAPIMap takes any of the strings below
// as map keys, and returns the url for that endpoint
func BootDevAPIMap(URL string) (string, error) {
	api := make(map[string]string)
	// archmage data endpoint
	api["archmage"] = "https://api.boot.dev/v1/leaderboard_archmage"
	// general global stats endpoint
	api["stats"] = "https://api.boot.dev/v1/leaderboard_stats"
	// daily stats endpoint
	api["daily"] = "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"
	// discord / community karma endpoint
	api["karma"] = "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"
	// live course finish feed
	api["feed"] = "https://api.boot.dev/v1/lesson_completion_feed"
	// user data endpoint
	api["user"] = "https://api.boot.dev/v1/users"
	// user achievements endpoint
	api["achievements"] = "https://api.boot.dev/v1/users/achievements"
	// course progress endpoint
	api["progress"] = "https://api.boot.dev/v1/courses_progress"
	// boss progress, XP bonus, and user's XP contribution
	api["boss"] = "https://api.boot.dev/v1/boss_events_progress"
	// course progress for user
	api["courses_progress"] = "https://api.boot.dev/v1/courses_progress"
	// course list with UUIDs to match against user's course progress
	api["courses"] = "https://api.boot.dev/v1/courses"
	if _, ok := api[URL]; !ok {
		return "", errors.New("invalid api")

	}
	return api[URL], nil
}
