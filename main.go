package main

import (
	"fmt"
	"log"
	"mikevanwinkle/blog/handlers"
	"net/http"
)

func main() {
	fmt.Printf("Blah blah world %s", "test")

	http.HandleFunc("/", handlers.Home)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
