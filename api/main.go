package main

import (
	"log"
	"net/http"

	"food-introduction/src"
)

func main() {
	router := src.Router()

	log.Fatal(http.ListenAndServe(":8080", router))
}