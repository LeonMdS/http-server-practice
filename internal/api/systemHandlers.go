package api

import (
	"fmt"
	"net/http"
)

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}

func (cfg *APIConfig) metricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, err := fmt.Fprintf(w,
		`
<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited %d times!</p>
  </body>
</html>
	`,
		cfg.fileserverHits.Load())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error writing response", err)
	}
}
