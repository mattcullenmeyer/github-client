package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type github struct {
	Repository string `json:"Repository"`
	Stars      string `json:"Stars"`
}

// https://levelup.gitconnected.com/tutorial-how-to-create-a-cli-tool-in-golang-a0fd980264f
// https://www.geeksforgeeks.org/command-line-arguments-in-golang/

func getData(repo string) string {

	url := fmt.Sprintf("http://localhost:8080/api?repo=%s", repo)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	result := github{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		log.Fatal(err)
	}

	return result.Stars
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("At least one repository is required")
	}

	repos := os.Args[1:]

	for _, repo := range repos {
		stars := getData(repo)
		fmt.Printf("%s --> %s\n", repo, stars)
	}

}
