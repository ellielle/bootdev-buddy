package login

import (
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWriteKeys(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping writeKeys in short mode")
	}

	testToken := &BDToken{AccessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyVVVJRCI6IjcyNzllNjIzLWI5YjctNDEwZi04MTFmLTY4YzM4OTAxMzM2ZCIsImV4cCI6MTcxNjU2MjIyMSwiaWF0IjoxNzE2NTU4NjIxLCJpc3MiOiJib290LmRldiJ9.R-c2fTfhyUbYLHUGgtfMSPDBoRODA-P2FJ7guf4bDOU", RefreshToken: "YbMGXKsN2w8wvvZTIdYVwq7XNFpVuNDB"}

	cases := map[string]struct {
		input *BDToken
		want  error
	}{
		"no error": {input: testToken, want: nil},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := writeKeys(testToken)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatalf(diff)
			}
			contents, err := os.ReadFile(".bootdevbuddy.json")
			if err != nil {
				t.Fatal(err)
			}
			if ok := strings.Contains(string(contents), "access"); !ok {
				t.Fatal(contents)
			}

		})
	}
}

func TestReadKeys(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping readKeys in short mode")
	}

	cases := map[string]struct {
		want string
	}{
		"no error": {want: "YbMGXKsN2w8wvvZTIdYVwq7XNFpVuNDB"},
	}

	token, err := readKeys()
	if err != nil {
		t.Fatal(err)
	}

	for _, tc := range cases {
		got := token
		diff := cmp.Diff(got.RefreshToken, tc.want)
		if diff != "" {
			t.Fatal(diff)
		}
	}
}
