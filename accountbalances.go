package xrpldata

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// GetAccountBalancesOptions specifies the optional parameters for the
// GetAccountBalances method.
type GetAccountBalancesOptions struct {
	Date         time.Time
	Currency     string
	Counterparty string
	Limit        int
}

type accountBalancesResponse struct {
	Result   string
	Limit    int
	Balances []Balance
}

// GetAccountBalances returns all balances held or owed by a specific XRP Ledger account.
// https://xrpl.org/data-api.html#get-account-balances
func (c *Client) GetAccountBalances(address string, opts *GetAccountBalancesOptions) ([]Balance, *Response, error) {
	path := "/accounts/" + address + "/balances"

	if opts != nil {
		v := url.Values{}

		if !opts.Date.IsZero() {
			v.Set("date", FormatTime(opts.Date))
		}

		if opts.Currency != "" {
			v.Set("currency", opts.Currency)
		}

		if opts.Counterparty != "" {
			v.Set("counterparty", opts.Counterparty)
		}

		if opts.Limit != 0 {
			v.Set("limit", strconv.Itoa(opts.Limit))
		}

		path += "?" + v.Encode()
	}

	res, err := c.Do(http.MethodGet, path)

	data := &accountBalancesResponse{}
	err = json.Unmarshal(res.Body, data)
	if err != nil {
		return nil, res, err
	}

	return data.Balances, res, nil
}
