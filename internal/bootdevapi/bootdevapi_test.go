package bootdevapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TODO: add more tests as API develops

// Tests that the proper url is returned
// and that the function works itself
func TestMain(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"archmage": {input: "archmage", want: "https://api.boot.dev/v1/leaderboard_archmage"},
		"stats":    {input: "stats", want: "https://api.boot.dev/v1/leaderboard_stats"},
		"daily":    {input: "daily", want: "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"},
		"karma":    {input: "karma", want: "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"},
		"user":     {input: "user", want: "https://api.boot.dev/v1/users"},
		"feed":     {input: "feed", want: "https://api.boot.dev/v1/lesson_completion_feed"},
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
