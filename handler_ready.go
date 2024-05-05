package main

import "net/http"

func handlerReadiness(resWriter http.ResponseWriter, req *http.Request) {
	respondWithJSON(resWriter, 200, struct{}{})
}
