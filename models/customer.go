package models

type CustomerJSON struct {
	ID               string                      `json:"id,omitempty"`
	Email            string                      `json:"email,omitempty"`
	Name             string                      `json:"name,omitempty"`
	TaxExempt        string                      `json:"tax_exempt,omitempty"`
	PaymentMethod    string                      `json:"payment_method,omitempty"`
	InvoiceSettings  CustomerInvoiceSettingsJSON `json:"invoice_settings,omitempty"`
	Currency         string                      `json:"currency,omitempty"`
	PreferredLocales []string                    `json:"preferred_locales,omitempty"`
}
type CustomerInvoiceSettingsJSON struct {
	DefaultPaymentMethod string `json:"default_payment_method,omitempty"`
}
