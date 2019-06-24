package xrpldata

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetAccountOrdersOptions specifies the optional parameters for the
// GetAccountOrders method.
type GetAccountOrdersOptions struct {
	LedgerIndex int
	LedgerHash  string
	Date        time.Time
	Limit       int
	Marker      string
}

type accountOrdersResponse struct {
	Result string
	Orders []Order
}

// GetAccountOrders fetches standing orders in the order books, placed by a
// specific account.
// https://xrpl.org/data-api.html#get-account-orders
func (c *Client) GetAccountOrders(address string, opts *GetAccountOrdersOptions) ([]Order, *Response, error) {
	path := "/accounts/" + address + "/orders"

	if opts != nil {
		v := url.Values{}

		if !opts.Date.IsZero() {
			v.Set("start", FormatTime(opts.Date))
		}

		if opts.Limit != 0 {
			v.Set("limit", strconv.Itoa(opts.Limit))
		}

		path += "?" + v.Encode()
	}

	res, err := c.Do(http.MethodGet, path)

	data := &accountOrdersResponse{}
	err = json.Unmarshal(res.Body, data)
	if err != nil {
		return nil, res, err
	}

	return data.Orders, res, nil
}
