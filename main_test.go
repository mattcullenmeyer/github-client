package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {

	tt := []struct {
		name string
		repo string
		want string
	}{
		{
			name: "valid repo",
			repo: "mattcullenmeyer/anaplan",
			want: "404",
		},
	}
	// Loop through each subtest
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			// Simulate passing arguments
			// https://stackoverflow.com/questions/33723300/how-to-test-the-passing-of-arguments-in-golang
			os.Args = []string{"cmd", tc.repo}

			main()

		})
	}
}
