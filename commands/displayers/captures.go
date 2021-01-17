package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieCapturesList wrapper for displaying
type MollieCapturesList struct {
	*mollie.CapturesList
}

// KV is a displayable group of key value.
func (cl *MollieCapturesList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, c := range cl.Embedded.Captures {
		x := buildXCapture(c)

		out = append(out, x)
	}

	return out
}

// MollieCapture wrapper for displaying.
type MollieCapture struct {
	*mollie.Capture
}

// KV is a displayable group of key value.
func (c *MollieCapture) KV() []map[string]interface{} {
	var out []map[string]interface{}
	x := buildXCapture(c.Capture)

	out = append(out, x)

	return out
}

func buildXCapture(c *mollie.Capture) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":          c.Resource,
		"ID":                c.ID,
		"MODE":              fallbackSafeMode(c.Mode),
		"AMOUNT":            fallbackSafeAmount(c.Amount),
		"SETTLEMENT_AMOUNT": fallbackSafeAmount(c.SettlementAmount),
		"PAYMENT_ID":        c.PaymentID,
		"SHIPMENT_ID":       c.ShipmentID,
		"SETTLEMENT_ID":     c.SettlementID,
		"CREATED_AT":        fallbackSafeDate(c.CreatedAt),
	}
}
