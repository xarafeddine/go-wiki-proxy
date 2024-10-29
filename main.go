package main

import (
	"log"
	"net/http"
)

// Server Port
const PORT = "4000"

func main() {
	// Register the handler for all paths
	http.HandleFunc("/", proxyHandler)

	// Start the server
	serverAddr := ":" + PORT
	log.Printf("Starting proxy server on http://localhost%s", serverAddr)

	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
