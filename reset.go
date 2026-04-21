package main

import "net/http"

func (cfg *apiConfig)handlerReset(w http.ResponseWriter, req *http.Request) {
	cfg.fileserverHits.Swap(0)
	w.WriteHeader(200)
	w.Write([]byte("Stats reset"))
}
