package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	// API route
	http.HandleFunc("/api/hello", helloHandler)

	// Frontend static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Logging
	log.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Go",
	})
}