package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

type AuthInfo struct {
	Key  string `json:"key"`
	ID   string `json:"id"`
	Info string `json:"info"`
}

type AuthReponse struct {
	Ok       bool     `json:"ok"`
	AuthInfo AuthInfo `json:"authinfo"`
}

var (
	keyStore     map[string]AuthInfo
	keyStoreLock sync.RWMutex
)

func AddKey(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var authInfo AuthInfo
	err := json.NewDecoder(r.Body).Decode(&authInfo)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	keyStoreLock.Lock()
	keyStore[authInfo.Key] = authInfo
	keyStoreLock.Unlock()
}

func ValidateKey(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key := ps.ByName("key")

	keyStoreLock.RLock()
	authInfo, exists := keyStore[key]
	keyStoreLock.RUnlock()

	authResponse := AuthReponse{
		Ok:       exists,
		AuthInfo: authInfo,
	}

	jsonString, err := json.Marshal(authResponse)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(jsonString))
}

func main() {
	keyStore = make(map[string]AuthInfo)
	router := httprouter.New()
	router.POST("/addKey", AddKey)
	router.GET("/validateKey/:key", ValidateKey)

	log.Fatal(http.ListenAndServe(":8080", router))
}
