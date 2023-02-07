package main

import (
	"encoding/json"
	"net/http"
)

const ResponseOk = "OK"

type healthResponse struct {
	Status string `json:"status" default:"OK"`
}

func newHealthResponse() healthResponse {
	return healthResponse{Status: ResponseOk}
}

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := newHealthResponse()
	json.NewEncoder(w).Encode(response)

	app.logger.Printf("health check")
}
