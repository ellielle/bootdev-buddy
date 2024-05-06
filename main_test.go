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
		"leaderboard": {input: "leaderboard", want: "https://api.boot.dev/v1/leaderboard_archmage"},
	}

	// Tests that the proper url is returned
	// and that the function works itself
	// TODO: expand as app grows
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := bootDevAPI(tc.input)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
