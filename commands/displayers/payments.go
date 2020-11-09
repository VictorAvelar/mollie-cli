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
		ped := getSafeExpiration(p)
		m := getSafePaymentMethod(p)

		x := map[string]interface{}{
			"ID":          p.ID,
			"Mode":        p.Mode,
			"Created":     p.CreatedAt.Format("02-01-2006"),
			"Expires":     ped,
			"Cancelable":  p.IsCancellable,
			"Amount":      p.Amount.Value + " " + p.Amount.Currency,
			"Method":      m,
			"Description": p.Description,
		}

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
	ped := getSafeExpiration(*p.Payment)
	m := getSafePaymentMethod(*p.Payment)
	x := map[string]interface{}{
		"ID":          p.ID,
		"Mode":        p.Mode,
		"Created":     p.CreatedAt.Format("02-01-2006"),
		"Expires":     ped,
		"Cancelable":  p.IsCancellable,
		"Amount":      p.Amount.Value + " " + p.Amount.Currency,
		"Method":      m,
		"Description": p.Description,
	}
	out = append(out, x)
	return out
}

func getSafeExpiration(p mollie.Payment) string {
	if p.ExpiresAt.IsZero() {
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
