package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"status":"healthy"}`)
	})

	http.HandleFunc("/api/v1/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message":"Platform API v1"}`)
	})

	log.Printf("Starting platform-api on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
