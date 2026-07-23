// Package api deals with internal api logic
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func respondWithJSON(w http.ResponseWriter, code int, response any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error encoding the response", err)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{Error: message})
}

func chirpCleaner(body string) string {
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	wordList := strings.Fields(body)

	for i, word := range wordList {
		if _, ok := badWords[strings.ToLower(word)]; ok {
			wordList[i] = strings.Repeat("*", 4)
		}
	}

	return strings.Join(wordList, " ")
}
