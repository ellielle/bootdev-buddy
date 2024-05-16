package cache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key  string
		want []byte
	}{
		{
			key:  "https://api.boot.dev",
			want: []byte(`{"IsAdmin": false}`),
		},
		{
			key:  "https://api.boot.dev/v1/leaderboard_archmage",
			want: []byte(`{"FirstName": Noelle}`),
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, &tc.want)
			val, ok := cache.Get(tc.key)
			if !ok {
				t.Fatalf("expected to find key")
				return
			}
			if string(val) != string(tc.want) {
				t.Fatalf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const interval = 5 * time.Millisecond
	const waitTime = interval + 5*time.Millisecond

	cases := []struct {
		key  string
		want []byte
	}{
		{
			key: "https://api.boot.dev/",
		},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			testData := []byte("testdata")
			cache.Add(tc.key, &testData)

			_, ok := cache.Get("https://api.boot.dev/")
			if !ok {
				t.Fatalf("expected to find key")
				return
			}
			time.Sleep(waitTime)

			_, ok = cache.Get("https://api.boot.dev/")
			if ok {
				t.Fatalf("expected to not find key")
				return
			}
		})
	}
}
