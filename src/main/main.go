package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.HandleFunc("/duration", api.ApiCalculateDate)
	http.Handle("/web/", http.StripPrefix("/web", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
