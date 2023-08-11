package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GetBear001(clientID, clientSecret string) string {
	///clientID := "f45ec821-7469-4a52-9e77-33df6577a98e"
	//clientSecret := "-6u1fCCGZ-UnUagPMNaAIzUAsW8fYlwo"

	// Construct the request body
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "system/Observation.read system/Patient.read")
	body := data.Encode()

	// Create the request
	url := "https://authorization.cerner.com/tenants/ec2458f2-1e24-41c8-b71b-0e701af7583d/protocols/oauth2/profiles/smart-v1/token"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "nothing 00"
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Encode the client credentials using base64
	auth := clientID + ":" + clientSecret
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Set("Authorization", "Basic "+encodedAuth)

	// Set the request body
	req.Body = ioutil.NopCloser(strings.NewReader(body))

	// Perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return "nothing 01"
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "nothing 02"
	}

	// Print the response
	fmt.Println(string(respBody))

	return string(respBody)
}

func handleGetBearerToken(w http.ResponseWriter, r *http.Request) {
	clientID := "f45ec821-7469-4a52-9e77-33df6577a98e"
	clientSecret := "-6u1fCCGZ-UnUagPMNaAIzUAsW8fYlwo"

	bearerToken := GetBear001(clientID, clientSecret)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(bearerToken))
}

func handleAuthorize(w http.ResponseWriter, r *http.Request) {
	// ... ( handleAuthorize function)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// ... (handleCallback function)
}

func main() {
	http.HandleFunc("/get-bearer-token", handleGetBearerToken)
	http.HandleFunc("/authorize", handleAuthorize)
	http.HandleFunc("/callback", handleCallback)

	http.ListenAndServe(":8080", nil)
}
