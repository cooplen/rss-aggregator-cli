package main

import (
	"fmt"
	"net/http"

	"github.com/0atme41/rssagg/internal/auth"
	"github.com/0atme41/rssagg/internal/database"
)

// an http handler similar to a normal handler except with an authenticated
// user already passed in
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// because the above function doesn't match the http handler function signature
// this is a method that takes in an auth handler and turns it into a normal handler
func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
