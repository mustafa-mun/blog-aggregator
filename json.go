package main

import (
	"net/http"
	"encoding/json"
)


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
	w.WriteHeader(500)
	return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

func respondWithError(w http.ResponseWriter, code int, errorStr string) {
	type returnVals struct {
		Error string `json:"error"`
	}
	respBody := returnVals{
		Error: errorStr,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
	w.WriteHeader(500)
	return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}