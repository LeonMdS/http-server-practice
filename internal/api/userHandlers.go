package api

import (
	"encoding/json"
	"net/http"

	"github.com/LeonMdS/http-server-practice/internal/auth"
	"github.com/LeonMdS/http-server-practice/internal/database"
)

func (cfg *APIConfig) createUserHandler(w http.ResponseWriter, r *http.Request) {
	type req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	reqJSON := req{}
	if err := decoder.Decode(&reqJSON); err != nil {
		respondWithError(w, http.StatusBadRequest, "Something went wrong", err)
		return
	}

	hashedPassword, err := auth.HashPassword(reqJSON.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error hashing password", err)
	}

	creationParams := database.CreateUserParams{
		Email:          reqJSON.Email,
		HashedPassword: hashedPassword,
	}

	createdUser, err := cfg.db.CreateUser(r.Context(), creationParams)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}
	respondWithJSON(w, http.StatusCreated, createdUser)
}

func (cfg *APIConfig) resetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Not allowed", nil)
		return
	}
	cfg.fileserverHits.Store(0)
	err := cfg.db.DeleteAllUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}
