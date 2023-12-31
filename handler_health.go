package main

import "net/http"

func handlerHealth(w http.ResponseWriter, request *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
