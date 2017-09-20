package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome! I'm Puchalski a weather bot")
}

func WillItRainIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Todo show:")
}

func WillItRain(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    when := vars["when"]
    fmt.Fprintln(w, "Todo show:", when)
}
