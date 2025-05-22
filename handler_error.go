package main

import "net/http"

// handler that handles an error
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
