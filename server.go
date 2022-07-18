package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, err := json.MarshalIndent("Voting Service. Version 0.0.1", "\t", "\t")
	if err != nil {
		log.Panic(err)
	}
	_, err = w.Write(res)
	if err != nil {
		log.Panic(err)
	}
}

func server() {
	mux := http.NewServeMux()
	log.Println("Server started!")
	mux.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8000", mux))
}
