package main

import "net/http"

func handlerReady(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{ Response string }{Response: "Up and running"})
}
