package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func tempo() {

	localid := "116567"

	url := fmt.Sprintf("http://api.tempo.com/index.php?api_lang=br&localidad=%s&affiliate_id=%s", localid, Config.Settings.TempoToken)

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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	parserTempo(body)
}

func parserTempo(resp []byte) {

	type Interesting struct {
		Url string
	}

	type Forecast struct {
		DataSequence string `xml:"data_sequence,attr"`
		Value        string `xml:"value,attr"`
	}

	type Data struct {
		Forecasts []Forecast `xml:"forecast"`
	}

	type Variation struct {
		Name string `xml:"name"`
		Icon string `xml:"icon"`
		Data Data   `xml:"data"`
	}

	type Location struct {
		City string `xml:"city,attr"`
		Interesting
		Variations []Variation `xml:"var"`
	}

	type Result struct {
		XMLName xml.Name `xml:"report"`
		Loc     Location `xml:"location"`
	}

	var result = Result{}

	err := xml.Unmarshal([]byte(resp), &result)
	if err != nil {
		fmt.Printf("error: %result", err)
		return
	}

}
