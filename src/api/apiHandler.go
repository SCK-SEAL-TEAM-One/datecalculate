package api

import (
	"encoding/json"
	"net/http"
)

type Duration struct {
	Days string `json:"days"`
}

func ApiCalculateDate(responseWriter http.ResponseWriter, request *http.Request) {
	mockDuration := Duration{
		Days: "152 days",
	}

	durationResponse, _ := json.Marshal(mockDuration)
	responseWriter.Write(durationResponse)
}
