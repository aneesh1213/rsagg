package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string){
	if code > 499 {
		log.Println("Respondgin with 5xx error: ", msg)
	}

	type errRepsone struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errRepsone{
		Error: msg,
	});
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload);

	if err != nil {
		log.Printf("failed to marshal the response:  %v" , payload);
		w.WriteHeader(500);
		return;
	}

	w.Header().Add("Content-Type", "application/json");
	w.WriteHeader(code)
	w.Write(data)
}