package main

import (
	"fmt"

	"go.reizu.org/xrpldata"
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
	c := xrpldata.NewClient(nil)

	rate, res, err := c.GetExchangeRates(base, counter, nil)
	if err != nil {
		fmt.Println(err, res)
	}

	fmt.Printf("%s/%s = %s\n", base, counter, rate)
}

func printAccountBalances(address string) {
	c := xrpldata.NewClient(nil)

	opts := &xrpldata.GetAccountBalancesOptions{Limit: 10}
	exchanges, res, err := c.GetAccountBalances(address, opts)
	if err != nil {
		fmt.Println(err, res)
	}

	for _, ex := range exchanges {
		fmt.Println(ex)
	}
}

func printAccountExchanges(address string) {
	c := xrpldata.NewClient(nil)

	opts := &xrpldata.GetAccountExchangesOptions{Descending: true, Limit: 10}
	exchanges, res, err := c.GetAccountExchanges(address, opts)
	if err != nil {
		fmt.Println(err, res)
	}

	for _, ex := range exchanges {
		fmt.Println(ex)
	}
}

func printAccountOrders(address string) {
	c := xrpldata.NewClient(nil)

	opts := &xrpldata.GetAccountOrdersOptions{Limit: 5}
	orders, res, err := c.GetAccountOrders(address, opts)
	if err != nil {
		fmt.Println(err, res)
	}

	for _, o := range orders {
		fmt.Println(o)
	}
}

func normalize(amount float64) {
	c := xrpldata.NewClient(nil)

	opts := &xrpldata.NormalizeOptions{
		Currency:         "XRP",
		ExchangeCurrency: "USD",
		ExchangeIssuer:   "rhub8VRN55s94qWKDv6jmDy1pUykJzF3wq",
	}
	converted, rate, res, err := c.Normalize(amount, opts)
	if err != nil {
		fmt.Println(err, res)
	}

	fmt.Printf("Converted: %s, Rate: %s\n", converted, rate)
}

func main() {
	printRate(XRP, USDgx)
	// printRate(BTCgx, USDgx)
	// printRate(ETHgx, USDgx)
	// printRate(EURgx, USDgx)

	printAccountBalances(addr)

	printAccountExchanges(addr)

	printAccountOrders(addr)

	normalize(100)
}
