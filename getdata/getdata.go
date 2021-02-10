package getdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type github struct {
	Repository string `json:"Repository"`
	Stars      string `json:"Stars"`
	Message    string `json:"message"`
}

// Function to get data from stargazer server
func GetData(repo string) string {

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

	// Use Unmarshal to parse body of response to github struct
	result := github{}
	jsonErr := json.Unmarshal(body, &result)
	if jsonErr != nil {
		// Print error message of API if get request is invalid
		txtBody := ioutil.NopCloser(bytes.NewBuffer(body))
		log.Fatal(txtBody)
	}

	return result.Stars + result.Message
}
