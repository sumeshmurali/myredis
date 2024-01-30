package main

import (
	"log"
	"net/http"
)

func main() {
	router := GetRouter()
	log.Println("Listening on :4000")
	http.ListenAndServe(":4000", router)
}
