package main

import (
	"net/http"
	"book-api/router"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":8080", r)
}