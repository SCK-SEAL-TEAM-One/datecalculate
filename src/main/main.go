package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/duration", api.ApiCalculateDate)
	log.Println("Listening on port 3000")
	http.ListenAndServe(":3000", nil)
}
