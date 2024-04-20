package main

import (
	"log"
	"mikevanwinkle/blog/handlers"
	"net/http"
)

func main() {

	http.HandleFunc("GET /", handlers.Home)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
