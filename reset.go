package main

import (
	"context"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	if cfg.platform != "dev" {
		w.WriteHeader(403)
		w.Write([]byte("forbidden"))
		return
	}
	cfg.fileserverHits.Swap(0)
	err := cfg.queries.ResetUsers(context.Background())
	if err != nil {
		respondWithError(w, 500, "error resetting", err)
	}
	w.WriteHeader(200)
	w.Write([]byte("Stats reset"))
}
