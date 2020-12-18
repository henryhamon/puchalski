package main

import (
	"github.com/jinzhu/configor"
	"log"
	"net/http"
)

// Config is loaded from configuration file
var Config = struct {
	Settings struct {
		TempoToken string `default:""`
	}
}{}

func main() {
	configor.Load(&Config, "config.yml")
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
