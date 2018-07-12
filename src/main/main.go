package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/date/duration", api.ApiCalculateDate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
