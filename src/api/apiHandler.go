package api

import (
	"encoding/json"
	"net/http"
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
	mockResponse := DurationResponse{
		From:         "Thursday, 4 January 2018",
		To:           "Monday, 4 June 2018",
		Days:         "152 days",
		Years:        "5 months, 1 day",
		Seconds:      "13,132,800 seconds",
		Minutes:      "218,880 minutes",
		Hours:        "3,648 hours",
		Weeks:        "21 weeks and 5 days",
		RatioOfYears: "41.64% of 2018",
	}
	durationResponse, err := json.Marshal(mockResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(durationResponse)
}
