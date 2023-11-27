package handler

import (
	"encoding/json"
	"fmt"
	"monban/internal/apikey/repo"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func GetAPIKeysHandler(rw http.ResponseWriter, r *http.Request, tx pgx.Tx) {
	repo := repo.NewAuthRepo()
	apikeys, err := repo.GetAllAPIKeys(tx)
	if err != nil {
		http.Error(rw, "", http.StatusInternalServerError)
	}

	fmt.Println("HERE IS API KEYS", apikeys)

	jsonResponse, err := json.Marshal(apikeys)
	if err != nil {
		http.Error(rw, "failed to marshal", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write(jsonResponse)
}
