package displayers

import (
	"github.com/VictorAvelar/mollie-api-go/v3/mollie"
)

// MollieListMethods wrapper for displaying.
type MollieListMethods struct {
	*mollie.PaymentMethodsList
}

// KV is a displayable group of key value.
func (mlm *MollieListMethods) KV() []map[string]interface{} {
	out := outPrealloc(len(mlm.Embedded.Methods))

	for _, pm := range mlm.Embedded.Methods {
		x := buildXMethod(pm)

		out = append(out, x)
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (mlm *MollieListMethods) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"ISSUERS",
		"MIN_AMOUNT",
		"MAX_AMOUNT",
		"LOGO",
	}
}

// ColMap returns a list of columns and its description.
func (mlm *MollieListMethods) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":    "the resource name",
		"ID":          "the resource id",
		"DESCRIPTION": "the method description",
		"ISSUERS":     "the count of issuers for the payment method (when embed)",
		"MIN_AMOUNT":  "the min. amount supported by the payment method",
		"MAX_AMOUNT":  "the max. amount supported by the payment method",
		"LOGO":        "the payment method logo (1x)",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mlm *MollieListMethods) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mlm *MollieListMethods) Filterable() bool {
	return true
}

// MollieMethod wrapper for displaying.
type MollieMethod struct {
	*mollie.PaymentMethodDetails
}

// KV is a displayable group of key value.
func (pm *MollieMethod) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXMethod(pm.PaymentMethodDetails)

	out = append(out, x)

	return out
}

// Cols returns an array of columns available for displaying.
func (pm *MollieMethod) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"DESCRIPTION",
		"ISSUERS",
		"MIN_AMOUNT",
		"MAX_AMOUNT",
		"LOGO",
	}
}

// ColMap returns a list of columns and its description.
func (pm *MollieMethod) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":    "the resource name as specified by mollie",
		"ID":          "the payment method id",
		"DESCRIPTION": "the payment method description",
		"ISSUERS":     "the count of issuers for the payment method (when embed)",
		"MIN_AMOUNT":  "the min. amount supported by the payment method",
		"MAX_AMOUNT":  "the max. amount supported by the payment method",
		"LOGO":        "the payment method logo (1x)",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (pm *MollieMethod) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (pm *MollieMethod) Filterable() bool {
	return true
}

func buildXMethod(m *mollie.PaymentMethodDetails) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":    m.Resource,
		"ID":          m.ID,
		"DESCRIPTION": m.Description,
		"ISSUERS":     fallbackSafeIssuers(m.Issuers),
		"MIN_AMOUNT":  fallbackSafeAmount(m.MinimumAmount),
		"MAX_AMOUNT":  fallbackSafeAmount(m.MaximumAmount),
		"LOGO":        m.Image.Size1x,
	}
}
