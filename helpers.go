package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

const (
	// Base URLs for Wikipedia and our modified version
	WIKIPEDIA_BASE_URL = "https://wikipedia.org"
	MODIFIED_BASE_URL  = "https://m-wikipedia.org"
)

// Bonus :) Custom HTTP client with reasonable timeouts
var client = &http.Client{
	Timeout: 30 * time.Second,
}

// fetchWikipediaContent retrieves content from Wikipedia for the given path
func fetchWikipediaContent(path string) ([]byte, error) {
	// Construct the full Wikipedia URL
	wikipediaURL := WIKIPEDIA_BASE_URL + path

	// Create and execute the request
	req, err := http.NewRequest("GET", wikipediaURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add headers to make the request look more like a regular browser
	req.Header.Set("User-Agent", "Mozilla/5.0 WikiProxy/1.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml")

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error fetching content: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	return body, nil
}

// modifyWikipediaContent replaces Wikipedia links with m-wikipedia.org links
func modifyWikipediaContent(content []byte) []byte {
	// Regular expressions for different types of Wikipedia links
	patterns := []string{
		`href="https?://[^/]*wikipedia\.org([^"]*)"`,
		`href="//[^/]*wikipedia\.org([^"]*)"`,
		`href="/wiki/([^"]*)"`,
	}

	modified := content

	// Apply each pattern
	for _, pattern := range patterns {
		re := regexp.MustCompile(pattern)

		// Replace links based on their format
		modified = re.ReplaceAllFunc(modified, func(match []byte) []byte {
			link := string(match)
			switch {
			case strings.Contains(link, "//"):
				// Handle absolute URLs
				return []byte(strings.Replace(link, "wikipedia.org", "m-wikipedia.org", 1))
			case strings.HasPrefix(link, `href="/`):
				// Handle relative URLs
				return []byte(fmt.Sprintf(`href="https://m-wikipedia.org%s"`, link[6:]))
			default:
				return match
			}
		})
	}

	return modified
}
