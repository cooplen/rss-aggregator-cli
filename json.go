package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// function that takes an http response and responds with an error in JSON format
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Reponding with 500 level error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJson(w, code, errResponse{
		Error: msg,
	})
}

// function that takes an http response, status code, and payload and makes them into JSON format
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
