package xrplda

import (
	"net/http"
	"net/http/httptest"
)

// setupTesting initializes a testing environment. It returns an XRPL Data API
// client, connected to a test http server, and a mux mounted on that server.
func setupTesting() (*Client, *http.ServeMux, func()) {
	mux := http.NewServeMux()
	s := httptest.NewServer(mux)
	c := NewClient(nil)
	c.BaseURL = s.URL
	teardown := s.Close
	return c, mux, teardown
}
