package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(resWriter http.ResponseWriter, statusCode int, msg string) {
	if statusCode > 499 {
		log.Printf("Responding with 5XX error: %v", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(resWriter, statusCode, errResponse{
		Error: msg,
	})
}

func respondWithJSON(resWriter http.ResponseWriter, statusCode int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		resWriter.WriteHeader(500)
		return
	}

	resWriter.Header().Add("Content-Type", "application/json")
	resWriter.WriteHeader(statusCode)
	resWriter.Write(data)
}
