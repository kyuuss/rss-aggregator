package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(writer http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Println("Responding with Error:", message)
	}

	type errorResonse struct {
		Error string `json:"error"`
	}

	respondWithJSON(writer, code, errorResonse{
		Error: message,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal reponse %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}
