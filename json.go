package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	respondwithJSON(w, code, errResponse{Error: msg})
}
