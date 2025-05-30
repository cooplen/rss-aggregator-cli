package main

import "net/http"

// handler that handles an error
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Internal Server Error")
}
