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
		x := map[string]interface{}{
			"ID":          p.ID,
			"Mode":        p.Mode,
			"Created":     p.CreatedAt.String(),
			"Expires":     p.ExpiresAt.String(),
			"Cancelable":  p.IsCancellable,
			"Amount":      p.Amount.Value + p.Amount.Currency,
			"Method":      p.Method,
			"Description": p.Description,
		}

		out = append(out, x)
	}

	return out
}
