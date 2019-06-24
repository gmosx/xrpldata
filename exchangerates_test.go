package xrpldata

import (
	"fmt"
	"net/http"
	"testing"
)

func TestExchangeRates(t *testing.T) {
	c, mux, teardown := setupTesting()
	defer teardown()

	base, counter := "USD+rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q", "XRP"

	mux.HandleFunc("/v2/exchange_rates/"+base+"/"+counter, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"result": "success","rate": "224.65709"}`)
	})

	rate, _, err := c.GetExchangeRates(base, counter, nil)
	if err != nil {
		t.Errorf("GetExchangeRates returned error: %v", err)
	}

	want := "224.65709"
	if rate != want {
		t.Errorf("GetExchangeRates returned %s, want %s", rate, want)
	}
}
