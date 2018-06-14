package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"datecalculate"
)

type DurationResponse struct {
	From         string `json:"from"`
	To           string `json:"to"`
	Days         string `json:"days"`
	Years        string `json:"years"`
	Seconds      string `json:"seconds"`
	Minutes      string `json:"minutes"`
	Hours        string `json:"hours"`
	Weeks        string `json:"weeks"`
	RatioOfYears string `json:"ratioOfYears"`
}

func ApiCalculateDate(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	startDay, _ := strconv.Atoi(queryString.Get("startDay"))
	startMonth, _ := strconv.Atoi(queryString.Get("startMonth"))
	startYear, _ := strconv.Atoi(queryString.Get("startYear"))
	endDay, _ := strconv.Atoi(queryString.Get("endDay"))
	endMonth, _ := strconv.Atoi(queryString.Get("endMonth"))
	endYear, _ := strconv.Atoi(queryString.Get("endYear"))
	startDate := datecalculate.NewDate(startDay, startMonth, startYear)
	endDate := datecalculate.NewDate(endDay, endMonth, endYear)
	mockResponse := datecalculate.MakeJson(startDate, endDate)
	durationResponse, err := json.Marshal(mockResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(durationResponse)
}
