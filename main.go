package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT") // Get the port from the environment variable
	if port == "" {
		port = "8080" // Default to 8080 if not set
	}

	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on port %s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Cloud Run, Trigger funktioniert! Panda ")
}
