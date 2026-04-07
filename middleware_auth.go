package main

import (
	"fmt"
	"net/http"

	"github.com/Aryan9inja/RSS_Aggregator/internal/auth"
	"github.com/Aryan9inja/RSS_Aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403,
				fmt.Sprintln("Auth error:", err))
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400,
				fmt.Sprintln("Couldn't get user:", err))
		}

		handler(w, r, user)
	}
}
