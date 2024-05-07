package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMain(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"archmage": {input: "archmage", want: "https://api.boot.dev/v1/leaderboard_archmage"},
		"stats":    {input: "stats", want: "https://api.boot.dev/v1/leaderboard_stats"},
		"daily":    {input: "daily", want: "https://api.boot.dev/v1/leaderboard_xp/day?limit=30"},
		"karma":    {input: "karma", want: "https://api.boot.dev/v1/leaderboard_karma/alltime?limit=30"},
	}

	// Tests that the proper url is returned
	// and that the function works itself
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := bootDevAPI(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
