package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle_greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!!!")
}

func main() {
	http.HandleFunc("/greet", handle_greet)
	log.Fatal("Error starting server @ localhost:8001", http.ListenAndServe("localhost:8001", nil))
}
