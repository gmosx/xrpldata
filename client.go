package xrplda

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const rippleDataAPIV2BaseURL = "https://data.ripple.com"

// Response is a XRPL Data API response.
type Response struct {
	*http.Response
	Body []byte
}

// Client is a client for the XRPL Data API.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient returns a new XRPL Data API client initialized with default
// parameters.
func NewClient(httpClient *http.Client) *Client {
	hc := httpClient
	if hc == nil {
		hc = &http.Client{
			Timeout: time.Second * 10,
		}
	}

	return &Client{
		BaseURL:    rippleDataAPIV2BaseURL,
		HTTPClient: hc,
	}
}

// Do performs an HTTP request to an API endpoint.
func (c *Client) Do(meth, path string) (*Response, error) {
	u := fmt.Sprintf("%s/v2%s", c.BaseURL, path)
	hreq, err := http.NewRequest(meth, u, nil)
	if err != nil {
		return nil, err
	}

	hres, err := c.HTTPClient.Do(hreq)
	if err != nil {
		return nil, err
	}

	// TODO: check status code!

	body, err := ioutil.ReadAll(hres.Body)
	if err != nil {
		return nil, err
	}

	res := Response{
		hres,
		body,
	}

	return &res, nil
}
