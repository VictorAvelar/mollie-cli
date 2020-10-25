package displayers

import (
	"github.com/VictorAvelar/mollie-api-go/mollie"
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
		x := map[string]interface{}{
			"ID":          p.ID,
			"Mode":        p.Mode,
			"Created":     p.CreatedAt.Format("01-02-2006"),
			"Expires":     ped,
			"Cancelable":  p.IsCancellable,
			"Amount":      p.Amount.Value + " " + p.Amount.Currency,
			"Method":      p.Method,
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
	x := map[string]interface{}{
		"ID":          p.ID,
		"Mode":        p.Mode,
		"Created":     p.CreatedAt.Format("02-01-2006"),
		"Expires":     p.ExpiresAt.Format("02-01-2006"),
		"Cancelable":  p.IsCancellable,
		"Amount":      p.Amount.Value + " " + p.Amount.Currency,
		"Method":      p.Method,
		"Description": p.Description,
	}
	out = append(out, x)
	return out
}

func getSafeExpiration(p mollie.Payment) string {
	if p.ExpiresAt != nil {
		return p.ExpiresAt.Format("01-02-2006")
	}

	return ""
}
