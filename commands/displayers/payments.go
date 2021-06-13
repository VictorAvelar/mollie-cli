package displayers

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

// MollieListPayments wrapper for displaying.
type MollieListPayments struct {
	*mollie.PaymentList
}

// KV is a displayable group of key value.
func (lp *MollieListPayments) KV() []map[string]interface{} {
	out := outPrealloc(len(lp.Embedded.Payments))

	for i := range lp.Embedded.Payments {
		out = append(out, buildXPayment(&lp.Embedded.Payments[i]))
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (lp *MollieListPayments) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"STATUS",
		"CANCELABLE",
		"AMOUNT",
		"METHOD",
		"DESCRIPTION",
		"SEQUENCE",
		"REMAINING",
		"REFUNDED",
		"CAPTURED",
		"SETTLEMENT",
		"APP_FEE",
		"CREATED_AT",
		"AUTHORIZED_AT",
		"EXPIRES",
		"PAID_AT",
		"FAILED_AT",
		"CANCELED_AT",
		"CUSTOMER_ID",
		"SETTLEMENT_ID",
		"MANDATE_ID",
		"SUBSCRIPTION_ID",
		"ORDER_ID",
		"REDIRECT",
		"WEBHOOK",
		"LOCALE",
		"COUNTRY",
	}
}

// ColMap returns a list of columns and its description.
func (lp *MollieListPayments) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":        "indicates the response contains a payment object",
		"ID":              "the identifier uniquely referring to this payment",
		"MODE":            "the mode used to create this payment",
		"STATUS":          "the payment status",
		"CANCELABLE":      "whether or not the payment can be canceled",
		"AMOUNT":          "the amount of the payment",
		"METHOD":          "the payment method used for this payment",
		"DESCRIPTION":     "a short description of the payment",
		"SEQUENCE":        "indicates which type of payment this is in a recurring sequence",
		"REMAINING":       "the remaining amount that can be refunded",
		"REFUNDED":        "the total amount that is already refunded",
		"CAPTURED":        "the total amount that is already captured for this payment",
		"SETTLEMENT":      "this optional field will contain the amount that will be settled to your account",
		"APP_FEE":         "the application fee, if the payment was created with one",
		"CREATED_AT":      "the payment’s date and time of creation",
		"AUTHORIZED_AT":   "the date and time the payment became authorized,",
		"EXPIRES":         "the date and time the payment will expire",
		"PAID_AT":         "the date and time the payment became paid",
		"FAILED_AT":       "the date and time the payment failed",
		"CANCELED_AT":     "the date and time the payment was canceled",
		"CUSTOMER_ID":     "if a customer was specified upon payment creation",
		"SETTLEMENT_ID":   "the identifier referring to the settlement this payment was settled with",
		"MANDATE_ID":      "if the payment is a first or recurring payment",
		"SUBSCRIPTION_ID": "the ID of the subscription that triggered the payment",
		"ORDER_ID":        "the ID of the subscription that triggered the payment",
		"REDIRECT":        "the URL your customer will be redirected to after completing or canceling the payment process",
		"WEBHOOK":         "the URL Mollie will call as soon an important status change takes place",
		"LOCALE":          "the customer’s locale",
		"COUNTRY":         "this optional field contains your customer’s ISO 3166-1 alpha-2 country code",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (lp *MollieListPayments) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (lp *MollieListPayments) Filterable() bool {
	return true
}

// MolliePayment wrapper for displaying.
type MolliePayment struct {
	*mollie.Payment
}

// KV is a displayable group of key value.
func (p *MolliePayment) KV() []map[string]interface{} {
	var out []map[string]interface{}
	x := buildXPayment(p.Payment)
	out = append(out, x)
	return out
}

// Cols returns an array of columns available for displaying.
func (p *MolliePayment) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"STATUS",
		"CANCELABLE",
		"AMOUNT",
		"METHOD",
		"DESCRIPTION",
		"SEQUENCE",
		"REMAINING",
		"REFUNDED",
		"CAPTURED",
		"SETTLEMENT",
		"APP_FEE",
		"CREATED_AT",
		"AUTHORIZED_AT",
		"EXPIRES",
		"PAID_AT",
		"FAILED_AT",
		"CANCELED_AT",
		"CUSTOMER_ID",
		"SETTLEMENT_ID",
		"MANDATE_ID",
		"SUBSCRIPTION_ID",
		"ORDER_ID",
		"REDIRECT",
		"WEBHOOK",
		"LOCALE",
		"COUNTRY",
	}
}

// ColMap returns a list of columns and its description.
func (p *MolliePayment) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":        "indicates the response contains a payment object",
		"ID":              "the identifier uniquely referring to this payment",
		"MODE":            "the mode used to create this payment",
		"STATUS":          "the payment status",
		"CANCELABLE":      "whether or not the payment can be canceled",
		"AMOUNT":          "the amount of the payment",
		"METHOD":          "the payment method used for this payment",
		"DESCRIPTION":     "a short description of the payment",
		"SEQUENCE":        "indicates which type of payment this is in a recurring sequence",
		"REMAINING":       "the remaining amount that can be refunded",
		"REFUNDED":        "the total amount that is already refunded",
		"CAPTURED":        "the total amount that is already captured for this payment",
		"SETTLEMENT":      "this optional field will contain the amount that will be settled to your account",
		"APP_FEE":         "the application fee, if the payment was created with one",
		"CREATED_AT":      "the payment’s date and time of creation",
		"AUTHORIZED_AT":   "the date and time the payment became authorized,",
		"EXPIRES":         "the date and time the payment will expire",
		"PAID_AT":         "the date and time the payment became paid",
		"FAILED_AT":       "the date and time the payment failed",
		"CANCELED_AT":     "the date and time the payment was canceled",
		"CUSTOMER_ID":     "if a customer was specified upon payment creation",
		"SETTLEMENT_ID":   "the identifier referring to the settlement this payment was settled with",
		"MANDATE_ID":      "if the payment is a first or recurring payment",
		"SUBSCRIPTION_ID": "the ID of the subscription that triggered the payment",
		"ORDER_ID":        "the ID of the subscription that triggered the payment",
		"REDIRECT":        "the URL your customer will be redirected to after completing or canceling the payment process",
		"WEBHOOK":         "the URL Mollie will call as soon an important status change takes place",
		"LOCALE":          "the customer’s locale",
		"COUNTRY":         "this optional field contains your customer’s ISO 3166-1 alpha-2 country code",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (p *MolliePayment) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (p *MolliePayment) Filterable() bool {
	return true
}

func buildXPayment(p *mollie.Payment) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":        p.Resource,
		"ID":              p.ID,
		"MODE":            fallbackSafeMode(p.Mode),
		"STATUS":          p.Status,
		"CANCELABLE":      p.IsCancellable,
		"AMOUNT":          fallbackSafeAmount(p.Amount),
		"METHOD":          fallbackSafePaymentMethod(p.Method),
		"DESCRIPTION":     p.Description,
		"SEQUENCE":        fallbackSafeSequence(p.SequenceType),
		"REMAINING":       fallbackSafeAmount(p.AmountRemaining),
		"REFUNDED":        fallbackSafeAmount(p.AmountRefunded),
		"CAPTURED":        fallbackSafeAmount(p.AmountCaptured),
		"SETTLEMENT":      fallbackSafeAmount(p.SettlementAmount),
		"APP_FEE":         fallbackSafeAppFee(p.ApplicationFee),
		"CREATED_AT":      fallbackSafeDate(p.CreatedAt),
		"AUTHORIZED_AT":   fallbackSafeDate(p.AuthorizedAt),
		"EXPIRES":         fallbackSafeDate(p.ExpiresAt),
		"PAID_AT":         fallbackSafeDate(p.PaidAt),
		"FAILED_AT":       fallbackSafeDate(p.FailedAt),
		"CANCELED_AT":     fallbackSafeDate(p.CanceledAt),
		"CUSTOMER_ID":     p.CustomerID,
		"SETTLEMENT_ID":   p.SettlementID,
		"MANDATE_ID":      p.MandateID,
		"SUBSCRIPTION_ID": p.SubscriptionID,
		"ORDER_ID":        p.OrderID,
		"REDIRECT":        p.RedirectURL,
		"WEBHOOK":         p.WebhookURL,
		"LOCALE":          fallbackSafeLocale(p.Locale),
		"COUNTRY":         p.CountryCode,
	}
}
