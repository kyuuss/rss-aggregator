package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respontWithJSON(w http.ResponseWriter, code string, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal reponse %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}
