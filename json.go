package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(
	responseWriter http.ResponseWriter,
	code int,
	payload any,
) {
	// Convert payload into json slice (encoded as byte)
	data, err := json.Marshal(payload)
	if err!=nil{
		log.Printf("Failed to Marshal JSON response: %v",payload)
		responseWriter.WriteHeader(500)
		return
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	responseWriter.WriteHeader(code)
	responseWriter.Write(data)
}

func respondWithError(
	responseWriter http.ResponseWriter,
	code int,
	msg string,
){
	if code > 499{
		log.Println("Responding with 500 error", msg)
	}

	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJSON(responseWriter, code, errResponse{
		Error: msg,
	})
}