package xrplda

// Exchange represents an actual exchange of currency, which can occur in the 
// XRP Ledger as the result of executing either an OfferCreate transaction or a 
// Payment transaction.
type Exchange struct {
	BaseAmount          string `json:"base_amount,omitempty"`
	BaseIssuer          string `json:"base_issuer,omitempty"`
	BaseCurrency        string `json:"base_currency,omitempty"`
	CounterAmount       string `json:"counter_amount,omitempty"`
	CounterIssuer       string `json:"counter_issuer,omitempty"`
	CounterCurrency     string `json:"counter_currency,omitempty"`
	AutobridgedCurrency string `json:"autobridged_currency,omitempty"`
	AutobridgedIssuer   string `json:"autobridged_issuer,omitempty"`
	Rate                string `json:"rate,omitempty"`
	Buyer               string `json:"buyer,omitempty"`
	Seller              string `json:"seller,omitempty"`
	Provider            string `json:"provider,omitempty"`
	Taker               string `json:"taker,omitempty"`
	ExecutedTime        string `json:"executed_time,omitempty"`
	OfferSequence       int    `json:"offer_sequence,omitempty"`
	TXHash              string `json:"tx_hash,omitempty"`
	TXType              string `json:"tx_type,omitempty"`
}
