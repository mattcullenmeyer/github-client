package main

import (
	"os"
	"strings"
	"testing"
)

func TestCliResponse(t *testing.T) {

	tt := []struct {
		name string
		args []string
		url  string
		want string
	}{
		{
			name: "no arguments",
			args: []string{},
			want: "At least one <organization>/<repository>",
		},
		{
			name: "invalid argument",
			args: []string{"abcdef"},
			want: "The format should be <organization>/<repository>",
		},
		{
			name: "invalid repo",
			args: []string{"abc/def"},
			want: "not a valid public GitHub repository",
		},
		{
			name: "one valid repo",
			args: []string{"mattcullenmeyer/anaplan"},
			want: "mattcullenmeyer/anaplan -->",
		},
		{
			name: "two valid repos",
			args: []string{"mattcullenmeyer/anaplan", "mattcullenmeyer/tinytrader"},
			want: "mattcullenmeyer/tinytrader -->",
		},
	}
	// Loop through each subtest
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// Simulate passing arguments
			// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang
			os.Args = tc.args

			txt := cliResponse(os.Args, "http://localhost:8080")

			if !strings.Contains(txt, tc.want) {
				t.Errorf("%s response doesn't include '%s' as expected; found %s", tc.name, tc.want, txt)
			}

		})
	}
}
