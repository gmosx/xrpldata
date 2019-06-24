package xrpldata

// Balance represents an account's balance in a specific currency with a
// specific counterparty at a single point in time.
type Balance struct {
	Counterparty string `json:"counterparty,omitempty"`
	Currency     string `json:"currency,omitempty"`
	Value        string `json:"value,omitempty"`
}

// OrderSpecification specifies an order's current state.
type OrderSpecification struct {
	Direction  string  `json:"direction,omitempty"`
	Quantity   Balance `json:"quantity,omitempty"`
	TotalPrice Balance `json:"total_price,omitempty"`
}

// OrderProperties specifies how an order was placed.
type OrderProperties struct {
	Maker             string `json:"maker,omitempty"`
	Sequence          int    `json:"sequence,omitempty"`
	MakerExchangeRate string `json:"maker_exchange_rate,omitempty"`
}

// Order represents a standing order.
type Order struct {
	Specification OrderSpecification
	Properties    OrderProperties
}

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
