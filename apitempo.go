package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func tempo() {

	id := "59ekgxjori64"
	localid := "116567"

	url := fmt.Sprintf("http://api.tempo.com/index.php?api_lang=br&localidad=%s&affiliate_id=%s", localid, id)

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()



}
