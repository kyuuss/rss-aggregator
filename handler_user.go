package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/kyuuss/rss-aggregator/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, request *http.Request) {
	type parametes struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(request.Body)

	params := parametes{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %s", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(request.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create user: %s", err))
	}

	respondWithJSON(w, 200, user)
}
