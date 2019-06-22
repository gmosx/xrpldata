package xrplda

import (
	"encoding/json"
	"net/http"
	"net/url"
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
	Format     string
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
