package displayers

import (
	"strings"

	"github.com/VictorAvelar/mollie-api-go/mollie"
)

type MollieListMethods struct {
	*mollie.ListMethods
}

func (mlm *MollieListMethods) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, pm := range mlm.Embedded.Methods {

		normalizePaymentMethodInfo(&pm)

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

	normalizePaymentMethodInfo(pm.PaymentMethodInfo)

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
	return strings.Join(vals, sep)
}

func normalizePaymentMethodInfo(pm *mollie.PaymentMethodInfo) *mollie.PaymentMethodInfo {
	if pm.MaximumAmount == nil && pm.MinimumAmount != nil {
		pm.MaximumAmount = &mollie.Amount{Currency: pm.MinimumAmount.Currency, Value: "N.A."}
	}
	if pm.MinimumAmount == nil && pm.MaximumAmount != nil {
		pm.MinimumAmount = &mollie.Amount{Currency: pm.MaximumAmount.Currency, Value: "N.A."}
	}

	if pm.MaximumAmount == nil && pm.MinimumAmount == nil {
		pm.MaximumAmount = &mollie.Amount{Currency: "N.A.", Value: "N.A."}
		pm.MinimumAmount = &mollie.Amount{Currency: "N.A.", Value: "N.A."}
	}

	return pm
}