package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func getRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return []byte{}
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte{}
	}

	return body

}
func postRequest(data string, url string) []byte {
	// URL to send the POST request to

	// Prepare POST data (JSON format)
	postData := fmt.Sprintf(`{
		"title": "foo",
		"body": "%s",
		"userId": 1
	}`, data)

	// Create a new POST request
	req, err := http.NewRequest("POST", url, strings.NewReader(postData))
	if err != nil {
		fmt.Println("Error creating POST request:", err)
		return []byte{}
	}

	// Set content type to application/json
	req.Header.Set("Content-Type", "application/json")

	// Send the POST request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return []byte{}
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return []byte{}
	}

	// Update the POST response label with the data
	return body

	// Get the window's label to update the response
}
