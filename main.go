package main

import (
	"fmt"
	"os"
)

// https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f

func getArgs() []string {
	if len(os.Args) < 2 {
		fmt.Println("At least one repository is required")
	}

	repos := os.Args[1:]

	return repos

}

func main() {

	repos := getArgs()

	for _, repo := range repos {
		fmt.Println(repo)
	}

}
