package main

import (
	"api"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))
	http.HandleFunc("/duration", api.ApiCalculateDate)

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
