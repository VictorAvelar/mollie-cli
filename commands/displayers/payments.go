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
		"APP FEE":         fallbackSafeAppFee(p.ApplicationFee),
		"CREATED AT":      fallbackSafeDate(p.CreatedAt),
		"AUTHORIZED AT":   fallbackSafeDate(p.AuthorizedAt),
		"EXPIRES":         fallbackSafeDate(p.ExpiresAt),
		"PAID AT":         fallbackSafeDate(p.PaidAt),
		"FAILED AT":       fallbackSafeDate(p.FailedAt),
		"CANCELED AT":     fallbackSafeDate(p.CanceledAt),
		"CUSTOMER ID":     p.CustomerID,
		"SETTLEMENT ID":   p.SettlementID,
		"MANDATE ID":      p.MandateID,
		"SUBSCRIPTION ID": p.SubscriptionID,
		"ORDER ID":        p.OrderID,
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
