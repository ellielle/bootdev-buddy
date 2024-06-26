package bootdevapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Tests that the proper url is returned
// and that the function works itself
func TestMain(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"archmage":         {input: "archmage", want: "https://api.boot.dev/v1/leaderboard_archmage"},
		"stats":            {input: "stats", want: "https://api.boot.dev/v1/leaderboard_stats"},
		"daily":            {input: "daily", want: "https://api.boot.dev/v1/leaderboard_xp/day?limit=10"},
		"karma":            {input: "karma", want: "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=10"},
		"user":             {input: "user", want: "https://api.boot.dev/v1/users"},
		"achievements":     {input: "achievements", want: "https://api.boot.dev/v1/users/achievements"},
		"feed":             {input: "feed", want: "https://api.boot.dev/v1/lesson_completion_feed"},
		"progress":         {input: "progress", want: "https://api.boot.dev/v1/courses_progress"},
		"boss":             {input: "boss", want: "https://api.boot.dev/v1/boss_events_progress"},
		"courses_progress": {input: "courses_progress", want: "https://api.boot.dev/v1/courses_progress"},
		"courses":          {input: "courses", want: "https://api.boot.dev/v1/courses"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := BootDevAPIMap(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
