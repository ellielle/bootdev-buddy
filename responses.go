package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondwithError(w http.ResponseWriter, code int, message string) error {
	type returnError struct {
		Error string `json:"error"`
	}
	return respondWithJSON(w, code, returnError{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return nil
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}
