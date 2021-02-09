package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type github struct {
	Repository string `json:"Repository"`
	Stars      string `json:"Stars"`
}

// https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f
// https://www.geeksforgeeks.org/command-line-arguments-in-golang/

func getData(repo string) string {

	// Check to make sure API server is running
	res, err := http.Get("http://localhost:8080/api")
	if err != nil {
		// Print that connection could not be made if API server is down
		log.Fatal(err)
	}

	url := fmt.Sprintf("http://localhost:8080/api?repo=%s", repo)

	res, err = http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Use Unmarshal to parse body or response to github struct
	result := github{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		// Print error message of API if get request is invalid
		txtBody := ioutil.NopCloser(bytes.NewBuffer(body))
		log.Fatal(txtBody)
	}

	return result.Stars
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("At least one <organization>/<repository> is required as an argument.")
		fmt.Println("For example:")
		fmt.Println("$ go run main.go mattcullenmeyer/tinytrader")
		fmt.Println("Or, if entering more than one:")
		fmt.Println("$ go run main.go mattcullenmeyer/tinytrader mattcullenmeyer/anaplan")
	}

	repos := os.Args[1:]

	for _, repo := range repos {
		matched, err := regexp.MatchString(`.\/.`, repo)
		if err != nil {
			log.Fatal(err)
		}

		if matched {
			stars := getData(repo)
			fmt.Printf("%s --> %s\n", repo, stars)
		} else {
			fmt.Printf("The argument '%s' is not valid. The format should be <organization>/<repository>.\n", repo)
			fmt.Printf("For example: mattcullenmeyer/tinytrader \n")
		}
	}

}
