package displayers

import (
	"strings"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

// MollieListMethods wrapper for displaying
type MollieListMethods struct {
	*mollie.ListMethods
}

// KV is a displayable group of key value
func (mlm *MollieListMethods) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, pm := range mlm.Embedded.Methods {
		ma := safeDisplayableAmount(pm.MinimumAmount)
		max := safeDisplayableAmount(pm.MaximumAmount)
		x := map[string]interface{}{
			"ID":             pm.ID,
			"Name":           pm.Description,
			"Minimum Amount": stringCombinator(" ", ma.Value, ma.Currency),
			"Maximum Amount": stringCombinator(" ", max.Value, max.Currency),
		}

		out = append(out, x)
	}

	return out
}

// MollieMethod wrapper for displaying
type MollieMethod struct {
	*mollie.PaymentMethodInfo
}

// KV is a displayable group of key value
func (pm *MollieMethod) KV() []map[string]interface{} {
	var out []map[string]interface{}

	ma := safeDisplayableAmount(pm.MinimumAmount)
	max := safeDisplayableAmount(pm.MaximumAmount)

	x := map[string]interface{}{
		"ID":             pm.ID,
		"Name":           pm.Description,
		"Minimum Amount": stringCombinator(" ", ma.Value, ma.Currency),
		"Maximum Amount": stringCombinator(" ", max.Value, max.Currency),
	}

	out = append(out, x)

	return out
}

func safeDisplayableAmount(a *mollie.Amount) *mollie.Amount {
	if a == nil {
		return &mollie.Amount{
			Currency: "---",
			Value:    "-----",
		}
	}

	return a
}

func stringCombinator(sep string, vals ...string) string {
	for i, v := range vals {
		if v == "" {
			vals[i] = "-"
		}
	}
	return strings.Join(vals, sep)
}
