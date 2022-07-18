package main

import (
	"log"
	"net/http"
)

func server() {
	mux := http.NewServeMux()
	log.Println("Server started!")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
