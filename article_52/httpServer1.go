package main

import (
	"fmt"
	"log"
	"net/http"
)

func dataHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, `"x": [1, 2, 3, 4, 5]`)
}

func otherHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, `foobar`)
}

func startHttpServer(address string) {
	log.Printf("Starting server on address %s", address)
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/data", dataHandler)
	http.HandleFunc("/other", otherHandler)
	http.ListenAndServe(address, nil)
}

func main() {
	startHttpServer(":8080")
}
