package xrpldata

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetAccountExchangesOptions specifies the optional parameters for the
// GetAccountExchanges method.
type GetAccountExchangesOptions struct {
	Start      time.Time
	End        time.Time
	Descending bool
	Limit      int
	Marker     string
}

type accountExchangesResponse struct {
	Result    string
	Count     int
	Marker    string
	Exchanges []Exchange
}

// GetAccountExchanges retrieve exchanges for a given account over time.
func (c *Client) GetAccountExchanges(address string, opts *GetAccountExchangesOptions) ([]Exchange, *Response, error) {
	path := "/accounts/" + address + "/exchanges"

	if opts != nil {
		v := url.Values{}

		if !opts.Start.IsZero() {
			v.Set("start", FormatTime(opts.Start))
		}

		if !opts.End.IsZero() {
			v.Set("end", FormatTime(opts.End))
		}

		if opts.Descending {
			v.Set("descending", "true")
		}

		if opts.Limit != 0 {
			v.Set("limit", strconv.Itoa(opts.Limit))
		}

		if opts.Marker != "" {
			v.Set("marker", opts.Marker)
		}

		path += "?" + v.Encode()
	}

	res, err := c.Do(http.MethodGet, path)

	data := &accountExchangesResponse{}
	err = json.Unmarshal(res.Body, data)
	if err != nil {
		return nil, res, err
	}

	return data.Exchanges, res, nil
}
