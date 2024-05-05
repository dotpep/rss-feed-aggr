package main

import "net/http"

func handlerError(resWriter http.ResponseWriter, req *http.Request) {
	respondWithError(resWriter, 400, "Something went wrong")
}
