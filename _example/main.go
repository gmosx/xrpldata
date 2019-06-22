package main

import (
	"fmt"

	"go.reizu.org/xrplda"
)

const (
	XRP   = "XRP"
	USDbx = "USD+rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
	USDgx = "USD+rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
	EURgx = "EUR+rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq"
	BTCgx = "BTC+rchGBxcD1A1C2tdxF6papQYZ8kjRKMYcL"
	ETHgx = "ETH+rcA8X3TVMST1n3CJeAdGk1RdRCHii7N2h"
)

const (
	addr = "rsyDrDi9Emy6vPU78qdxovmNpmj5Qh4NKw"
)

func printRate(base, counter string) {
	c := xrplda.NewClient(nil)
	rate, res, err := c.GetExchangeRates(base, counter, nil)
	if err != nil {
		fmt.Println(err, res)
	}
	fmt.Printf("%s/%s = %s\n", base, counter, rate)
}

func printAccountExchanges(address string) {
	c := xrplda.NewClient(nil)
	opts := &xrplda.GetAccountExchangesOptions{Descending: true, Limit: 10}
	exchanges, res, err := c.GetAccountExchanges(address, opts)
	if err != nil {
		fmt.Println(err, res)
	}
	for _, ex := range exchanges {
		fmt.Println(ex)
	}
}

func printAccountOrders(address string) {
	c := xrplda.NewClient(nil)
	opts := &xrplda.GetAccountOrdersOptions{Limit: 5}
	orders, res, err := c.GetAccountOrders(address, opts)
	if err != nil {
		fmt.Println(err, res)
	}
	for _, o := range orders {
		fmt.Println(o)
	}
}

func main() {
	// printRate(XRP, USDgx)
	// printRate(BTCgx, USDgx)
	// printRate(ETHgx, USDgx)
	// printRate(EURgx, USDgx)

	// printAccountExchanges(addr)

	printAccountOrders(addr)
}
