package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/arjun256900/go-product-search-api/internal/search"
)

// HealthHandler returns a simple OK status.
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

// SearchHandler wraps the search.Service into an http.HandlerFunc.
func SearchHandler(svc *search.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		q := r.URL.Query().Get("q")		// The user query (to be searched)

		// No query provided edge case
		if q == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "query parameter 'q' is required",
			})
			return
		}

		// Execute search (limit - top 50)
		results, err := svc.Search(q, 50)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		// Return a JSON array of products
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}
}