package api

import (
	"encoding/json"
	"net/http"
)

type Duration struct {
	Days string `json:"days"`
}

func ApiCalculateDate(w http.ResponseWriter, req *http.Request) {
	mockDuration := Duration{
		Days: "152 days",
	}

	durationResponse, _ := json.Marshal(mockDuration)
	w.Write(durationResponse)
}
