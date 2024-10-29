package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// proxyHandler processes incoming requests
func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Log incoming request
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	// Only handle GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fetch content from Wikipedia
	content, err := fetchWikipediaContent(r.URL.Path)
	if err != nil {
		log.Printf("Error fetching content: %v", err)
		http.Error(w, "Error fetching content", http.StatusInternalServerError)
		return
	}

	// Modify the content
	modifiedContent := modifyWikipediaContent(content)

	// Set appropriate headers
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprint(len(modifiedContent)))
	w.Header().Set("Server", "WikiProxy/1.0")

	// Write the modified content to the response
	if _, err := io.Copy(w, bytes.NewReader(modifiedContent)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
