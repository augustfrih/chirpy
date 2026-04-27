package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/augustfrih/chirpy/internal/database"
)

func (cfg *apiConfig) handlerInsertUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldnt decode parameters", err)
		return
	}
	var createdUser database.User
	createdUser, err = cfg.queries.CreateUser(context.Background(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "user already in database", err)
		return
	}

	jsonUser := User{
		ID:        createdUser.ID,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
		Email:     createdUser.Email,
	}

	respondWithJson(w, 201, jsonUser)
}
