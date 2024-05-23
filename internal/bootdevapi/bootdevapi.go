package bootdevapi

import "errors"

// BootDevAPIMap takes any of the strings below
// as map keys, and returns the url for that endpoint
// archmage = Archmage Leaderboard data
// stats = Global daily stats
// daily = Top daily learners - based on XP earned
// karma: Discord community leaderboard - based on a bunch of Discordiness
func BootDevAPIMap(URL string) (string, error) {
	api := make(map[string]string)
	api["archmage"] = "https://api.boot.dev/v1/leaderboard_archmage"
	api["stats"] = "https://api.boot.dev/v1/leaderboard_stats"
	api["daily"] = "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"
	api["karma"] = "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"
	api["feed"] = "https://api.boot.dev/v1/lesson_completion_feed"
	api["user"] = "https://api.boot.dev/v1/users"
	if _, ok := api[URL]; !ok {
		return "", errors.New("invalid api")

	}
	return api[URL], nil
}
