package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Handle static files
	if strings.HasPrefix(r.URL.Path, "/static/") {
		staticHandler(w, r)
		return
	}

	// Handle resource files from Wikipedia
	if strings.HasPrefix(r.URL.Path, "/w/") || strings.HasPrefix(r.URL.Path, "/static/") {
		content, contentType, err := fetchWikipediaContent(r.URL.Path)
		if err != nil {
			log.Printf("Error fetching resource: %v", err)
			http.Error(w, "Error fetching resource", http.StatusInternalServerError)
			return
		}

		// Set appropriate content type based on file extension
		if contentType == "" {
			ext := filepath.Ext(r.URL.Path)
			contentType = mime.TypeByExtension(ext)
		}

		w.Header().Set("Content-Type", contentType)
		w.Write(content)
		return
	}

	content, contentType, err := fetchWikipediaContent(r.URL.Path)
	if err != nil {
		log.Printf("Error fetching content: %v", err)
		http.Error(w, "Error fetching content", http.StatusInternalServerError)
		return
	}

	modifiedContent := modifyWikipediaContent(content, contentType)

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", fmt.Sprint(len(modifiedContent)))
	w.Header().Set("Server", "WikiProxy/1.0")

	if _, err := io.Copy(w, bytes.NewReader(modifiedContent)); err != nil {
		log.Printf("Error writing response: %v", err)
	}
}
