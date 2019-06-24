# XRPL Data

A client for the [Ripple Data API V2](https://xrpl.org/data-api.html). The Ripple Data API v2 provides access to information about changes in the XRP Ledger, including transaction history and processed analytical data.

## Example

```go
c := xrpldata.NewClient(nil)

rate, _, err := c.GetExchangeRates("XRP", "USD+rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q", nil)

opts := &xrpldata.GetAccountExchangesOptions{Descending: true, Limit: 10}
exchanges, _, err := c.GetAccountExchanges("rsyDrDi9Emy6vPU78qdxovmNpmj5Qh4NKw", opts)
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. Please make sure to update tests as appropriate.

## Contact

[@gmosx](https://twitter.com/gmosx) on Twitter.

## License

MIT, see [LICENSE](./LICENSE) file for details.

Copyright © 2019 George Moschovitis.
