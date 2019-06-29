package xrpldata

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// NormalizeOptions specifies the optional parameters for the
// Normalize method.
type NormalizeOptions struct {
	Currency         string
	Issuer           string
	ExchangeCurrency string
	ExchangeIssuer   string
	Date             time.Time
	Strict           bool
}

type normalizeResponse struct {
	Converted string
	Rate      string
}

// Normalize converts an amount from one currency and issuer to another, using
// the network exchange rates.
// https://xrpl.org/data-api.html#normalize
func (c *Client) Normalize(amount float64, opts *NormalizeOptions) (converted string, rate string, r *Response, err error) {
	path := "/normalize"

	v := url.Values{}

	v.Set("amount", strconv.FormatFloat(amount, 'f', -1, 64))

	if opts.Currency != "" {
		v.Set("currency", opts.Currency)
	}

	if opts.Issuer != "" {
		v.Set("issuer", opts.Issuer)
	}

	if opts.ExchangeCurrency != "" {
		v.Set("exchange_currency", opts.ExchangeCurrency)
	}

	if opts.ExchangeIssuer != "" {
		v.Set("exchange_issuer", opts.ExchangeIssuer)
	}

	if !opts.Date.IsZero() {
		v.Set("date", FormatTime(opts.Date))
	}

	if opts.Strict {
		v.Set("strict", "true")
	}

	path += "?" + v.Encode()

	r, err = c.Do(http.MethodGet, path)

	data := &normalizeResponse{}
	err = json.Unmarshal(r.Body, data)
	if err != nil {
		return "", "", r, err
	}

	return data.Converted, data.Rate, r, nil
}
