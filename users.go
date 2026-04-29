package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/augustfrih/chirpy/internal/auth"
	"github.com/augustfrih/chirpy/internal/database"
)

func (cfg *apiConfig) handlerInsertUser(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldnt decode parameters", err)
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldnt hash password", err)
	}

	userParams := database.CreateUserParams{
		Email: params.Email,
		HashedPassword: hashedPassword,
	}

	var createdUser database.User
	createdUser, err = cfg.queries.CreateUser(context.Background(), userParams)
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
