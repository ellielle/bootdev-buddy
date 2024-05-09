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

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, &c.want)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.want) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	testData := []byte("testdata")
	cache.Add("https://api.boot.dev/", &testData)

	_, ok := cache.Get("https://api.boot.dev/")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://api.boot.dev/")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
