package main

import (
	"encoding/json"
	"net/http"
)

func respondWithJson(w http.ResponseWriter, code int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}

func chirpValidationHandler(w http.ResponseWriter, r *http.Request) {
	type chirp struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	c := chirp{}
	err := decoder.Decode(&c)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong")
		return
	}
	if len(c.Body) > 140 {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	type response struct {
		Valid bool `json:"valid"`
	}
	resp := response{Valid: true}
	respondWithJson(w, http.StatusOK, resp)
}
