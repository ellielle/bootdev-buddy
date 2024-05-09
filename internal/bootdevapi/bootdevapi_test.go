package bootdevapi

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/ellielle/bootdev-stats/internal/cache"
)

func TestArchmageAPI(t *testing.T) {
	tests := map[string]struct {
		input string
		want  func()
	}{
		"GetArchmages": {input: "archmage", want: nil},
	}

	c := cache.NewCache(60 * time.Second)

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := GetArchmages("archmage", c)
			diff := cmp.Diff(tc.want, got)
			if strings.Contains(diff, "usupported protocol scheme") {
				t.Fatalf(diff)
			}
		})
	}
}
