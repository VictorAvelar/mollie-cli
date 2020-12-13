package displayers

import (
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
		x := buildXMethod(pm)

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

	x := buildXMethod(pm.PaymentMethodInfo)

	out = append(out, x)

	return out
}

func buildXMethod(m *mollie.PaymentMethodInfo) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":    m.Resource,
		"ID":          m.ID,
		"DESCRIPTION": m.Description,
		"MIN_AMOUNT":  fallbackSafeAmount(m.MinimumAmount),
		"MAX_AMOUNT":  fallbackSafeAmount(m.MaximumAmount),
		"LOGO":        m.Image.Size1x,
	}
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
