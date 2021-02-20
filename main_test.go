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
			url:  "http://192.168.49.2:30893",
		},
		{
			name: "invalid argument",
			args: []string{"abcdef"},
			want: "The format should be <organization>/<repository>",
			url:  "http://192.168.49.2:30893",
		},
		{
			name: "invalid repo",
			args: []string{"abc/def"},
			want: "not a valid public GitHub repository",
			url:  "http://192.168.49.2:30893",
		},
		{
			name: "one valid repo",
			args: []string{"mattcullenmeyer/anaplan"},
			want: "mattcullenmeyer/anaplan -->",
			url:  "http://192.168.49.2:30893",
		},
		{
			name: "two valid repos",
			args: []string{"mattcullenmeyer/anaplan", "mattcullenmeyer/tinytrader"},
			want: "mattcullenmeyer/tinytrader -->",
			url:  "http://192.168.49.2:30893",
		},
	}
	// Loop through each subtest
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// Simulate passing arguments
			// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang
			os.Args = tc.args
			//urlPtr := flag.String("url", "http://localhost:8080", "url path")
			//flag.Parse()
			//url := *urlPtr
			//repos := flag.Args()

			txt := cliResponse(os.Args, tc.url)

			if !strings.Contains(txt, tc.want) {
				t.Errorf("%s response doesn't include '%s' as expected; found %s", tc.name, tc.want, txt)
			}

		})
	}
}
