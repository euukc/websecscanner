package api

import (
	"encoding/json"
	"net/http"
	"websecscanner/internal/scanner"
)

// ScanHandler recebe uma URL e retorna os resultados do scanner
func ScanHandler(w http.ResponseWriter, r *http.Request) {
	type request struct {
		URL string `json:"url"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Request inválido", http.StatusBadRequest)
		return
	}

	results, err := scanner.ScanHeaders(req.URL)
	if err != nil {
		http.Error(w, "Erro no scan: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
