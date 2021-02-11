package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/mattcullenmeyer/github-stargazer-client/getdata"
)

// https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f
// https://www.geeksforgeeks.org/command-line-arguments-in-golang/

func getRepos(repos []string) string {

	// Initialize response as blank string
	response := ""

	// Loop through each command line argument
	for _, repo := range repos {

		// Check to make sure command line argument includes a "/"
		matched, err := regexp.MatchString(`.\/.`, repo)

		if err != nil {
			log.Fatal(err)
		}

		// Concatenate results from each repo
		if matched {
			stars := getdata.GetData(repo)
			response += fmt.Sprintf("%s --> %s\n", repo, stars)
		} else {
			response += fmt.Sprintf("The argument '%s' is not valid. The format should be <organization>/<repository>.\n", repo)
		}
	}
	return response
}

func cliResponse(args []string) string {

	n := len(args)

	if n < 2 {
		return `At least one <organization>/<repository> is required as an argument.
For example:
$ go run main.go mattcullenmeyer/tinytrader
Or, if entering more than one:
$ go run main.go mattcullenmeyer/tinytrader mattcullenmeyer/anaplan`
	}

	// Get all arguments afer the first, which is the name of the program itself
	repos := os.Args[1:]

	txt := getRepos(repos)
	return txt
}

func main() {

	txt := cliResponse(os.Args)
	fmt.Printf(txt)

}
