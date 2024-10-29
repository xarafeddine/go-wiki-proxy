# Wikipedia Proxy Server

A Go-based proxy server that serves Wikipedia content with modified URLs. This server fetches Wikipedia pages, replaces all Wikipedia URLs with an alternative domain (m-wikipedia.org), and serves the modified content directly to clients.


## Installation

### Prerequisites

- Go 1.16 or higher (I used go 1.23.1)
- Internet connection (to fetch Wikipedia content)

### Getting Started

1. Clone this repository:
```bash
git clone https://github.com/xarafeddine/go-wiki-proxy
cd go-wiki-proxy
```

2. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:4000`

## Usage

Access Wikipedia content through the proxy by using the same path structure as Wikipedia:

```
http://localhost:4000/wiki/Article_Name
```

For example:
- Original Wikipedia URL: `https://wikipedia.org/wiki/Go_(programming_language)`
- Proxy URL: `http://localhost:4000/wiki/Go_(programming_language)`

## Configuration

The following constants can be modified in `main.go`:

```go
const (
    PORT = "4000"                              // Server port
    WIKIPEDIA_BASE_URL = "https://wikipedia.org" // Source Wikipedia URL
    MODIFIED_BASE_URL = "https://m-wikipedia.org" // Target domain for modified links
)
```

## Example Response

Here's how the proxy modifies different types of Wikipedia URLs:

```html
<!-- Original Wikipedia HTML -->
<div class="content">
    <a href="https://wikipedia.org/wiki/Python">Python</a>
    <a href="//wikipedia.org/wiki/Java">Java</a>
    <a href="/wiki/Golang">Golang</a>
    <img src="//upload.wikipedia.org/wiki/images/logo.png">
</div>

<!-- Modified HTML served by proxy -->
<div class="content">
    <a href="https://m-wikipedia.org/wiki/Python">Python</a>
    <a href="https://m-wikipedia.org/wiki/Java">Java</a>
    <a href="https://m-wikipedia.org/wiki/Golang">Golang</a>
    <img src="//upload.wikipedia.org/wiki/images/logo.png">
</div>
```

### URL Modification Rules

The proxy handles various URL formats:

1. Absolute URLs:
   - Before: `href="https://wikipedia.org/wiki/Page"`
   - After: `href="https://m-wikipedia.org/wiki/Page"`

2. Protocol-relative URLs:
   - Before: `href="//wikipedia.org/wiki/Page"`
   - After: `href="https://m-wikipedia.org/wiki/Page"`

3. Relative URLs:
   - Before: `href="/wiki/Page"`
   - After: `href="https://m-wikipedia.org/wiki/Page"`

## Error Handling

The server handles various error cases:

- Invalid requests: Returns 405 Method Not Allowed
- Failed Wikipedia fetches: Returns 500 Internal Server Error
- Timeout issues: Returns 500 Internal Server Error with timeout message

All errors are logged with timestamps for debugging.

## Performance Considerations

- The server uses a custom HTTP client with a 30-second timeout
- Content is modified in-memory before being sent to the client
- Regular expressions are pre-compiled for better performance
- Response headers are properly set for browser caching
