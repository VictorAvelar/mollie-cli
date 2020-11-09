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
		x := map[string]interface{}{
			"ID":             pm.ID,
			"Name":           pm.Description,
			"Minimum Amount": stringCombinator("/", pm.MinimumAmount.Value, pm.MinimumAmount.Currency),
			"Maximum Amount": stringCombinator("/", pm.MaximumAmount.Value, pm.MaximumAmount.Currency),
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

	x := map[string]interface{}{
		"ID":             pm.ID,
		"Name":           pm.Description,
		"Minimum Amount": stringCombinator("/", pm.MinimumAmount.Value, pm.MinimumAmount.Currency),
		"Maximum Amount": stringCombinator("/", pm.MaximumAmount.Value, pm.MaximumAmount.Currency),
	}

	out = append(out, x)

	return out
}

func stringCombinator(sep string, vals ...string) string {
	for i, v := range vals {
		if v == "" {
			vals[i] = "-"
		}
	}
	return strings.Join(vals, sep)
}
