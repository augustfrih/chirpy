package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/augustfrih/chirpy/internal/database"
	"github.com/google/uuid"
)

type chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) handlerPostChirp(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserID uuid.UUID `json:"user_id"`
	}
	decoder := json.NewDecoder(req.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 500, "Error decoding parameters", err)
		return
	}

	if len(params.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long", nil)
		return
	}
	fmt.Println(params)
	// TODO: check that user is in users

	words := strings.Split(params.Body, " ")
	badWords := []string{"kerfuffle", "sharbert", "fornax"}

	for i, word := range words {
		for _, badWord := range badWords {
			if strings.ToLower(word) == badWord {
				words[i] = "****"
				break
			}
		}
	}

	fetchedChirp, err := cfg.queries.CreateChirp(context.Background(), database.CreateChirpParams{
		Body: strings.Join(words, " "),
		UserID: uuid.NullUUID{
			UUID:  params.UserID,
			Valid: true,
		},
	})
	if err != nil {
		respondWithError(w, 400, "Error creating chirp", err)
		return
	}

	respBody := chirp{
		ID:        fetchedChirp.ID,
		CreatedAt: fetchedChirp.CreatedAt,
		UpdatedAt: fetchedChirp.UpdatedAt,
		Body:      fetchedChirp.Body,
		UserID:    fetchedChirp.UserID.UUID,
	}

	respondWithJson(w, 201, respBody)
}

func (cfg *apiConfig) handlerGetChirps(w http.ResponseWriter, req *http.Request) {
	dbChirps, err := cfg.queries.GetChirps(context.Background())
	if err != nil {
		respondWithError(w, 500, "Error getting chirps", err)
		return
	}

	chirps := []chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			Body:      dbChirp.Body,
			UserID:    dbChirp.UserID.UUID,
		})
	}

	respondWithJson(w, 200, chirps)
}
