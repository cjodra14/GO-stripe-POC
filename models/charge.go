package models

type ChargeJSON struct {
	Amount       int64  `json:"amount,omitempty"`
	ReceiptEmail string `json:"receiptEmail,omitempty"`
	Currency     string `json:"currency,omitempty"`

}
