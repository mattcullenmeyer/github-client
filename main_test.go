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
		want string
	}{
		{
			name: "no arguments",
			args: []string{"cmd"},
			want: "At least one <organization>/<repository>",
		},
		{
			name: "invalid argument",
			args: []string{"cmd", "abcdef"},
			want: "The format should be <organization>/<repository>",
		},
		{
			name: "invalid repo",
			args: []string{"cmd", "abc/def"},
			want: "not a valid public GitHub repository",
		},
		{
			name: "one valid repo",
			args: []string{"cmd", "mattcullenmeyer/anaplan"},
			want: "mattcullenmeyer/anaplan -->",
		},
		{
			name: "two valid repos",
			args: []string{"cmd", "mattcullenmeyer/anaplan", "mattcullenmeyer/tinytrader"},
			want: "mattcullenmeyer/tinytrader -->",
		},
	}
	// Loop through each subtest
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// Simulate passing arguments
			// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang
			os.Args = tc.args

			txt := cliResponse(os.Args)

			if !strings.Contains(txt, tc.want) {
				t.Errorf("Response doesn't include '%s' as expected", tc.want)
			}

		})
	}
}
