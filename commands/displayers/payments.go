package displayers

import (
	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

// MollieListPayments wrapper for displaying.
type MollieListPayments struct {
	*mollie.PaymentList
}

// KV is a displayable group of key value
func (lp *MollieListPayments) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, p := range lp.Embedded.Payments {
		x := buildXPayment(&p)

		out = append(out, x)
	}

	return out
}

// MolliePayment wrapper for displaying.
type MolliePayment struct {
	*mollie.Payment
}

// KV is a displayable group of key value
func (p *MolliePayment) KV() []map[string]interface{} {
	var out []map[string]interface{}
	x := buildXPayment(p.Payment)
	out = append(out, x)
	return out
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

func getSafeExpiration(p mollie.Payment) string {
	if p.ExpiresAt == nil {
		return "----------"
	}

	return p.ExpiresAt.Format("02-01-2006")
}

func getSafePaymentMethod(p mollie.Payment) string {
	if p.Method == mollie.PaymentMethod("") {
		return "none"
	}
	return string(p.Method)
}
