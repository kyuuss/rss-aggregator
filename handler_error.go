package main

import "net/http"

func handlerError(w http.ResponseWriter, request *http.Request) {
	respondWithError(w, 200, "Something went wrong")
}
