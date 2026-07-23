package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/LeonMdS/chirpy-server/internal/auth"
	"github.com/google/uuid"
)

type loginParams struct {
	Password string
	Email    string
}

func (cfg *APIConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	requestParams := loginParams{}
	if err := json.NewDecoder(r.Body).Decode(&requestParams); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error decoding login request", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), requestParams.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	type userNoPassword struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	authorized, err := auth.CheckPasswordHash(requestParams.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	if authorized {
		response := userNoPassword{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		}
		respondWithJSON(w, http.StatusOK, response)
	} else {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
}
