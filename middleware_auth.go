package main

import (
	"net/http"
	"fmt"
	"github.com/aneesh1213/RssAgg-Go/internal/auth"
	"github.com/aneesh1213/RssAgg-Go/internal/database"
)

type authHandler func (http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request){
		apiKey, err := auth.GetApiKey(r.Header);

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err));
			return;
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey);
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt get the user Error: %v", err));
			return;
		}

		handler(w, r, user);
	}
}