package main

import "net/http"

// simple http handler that returns status 200 ok
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJson(w, 200, `"status": "ok"`)
}
