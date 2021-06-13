package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieCapturesList wrapper for displaying.
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

// Cols returns an array of columns available for displaying.
func (cl *MollieCapturesList) Cols() []string {
	return capturesCols()
}

// ColMap returns a list of columns and its description.
func (cl *MollieCapturesList) ColMap() map[string]string {
	return capturesColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (cl *MollieCapturesList) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (cl *MollieCapturesList) Filterable() bool {
	return true
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

// Cols returns an array of columns available for displaying.
func (c *MollieCapture) Cols() []string {
	return capturesCols()
}

// ColMap returns a list of columns and its description.
func (c *MollieCapture) ColMap() map[string]string {
	return capturesColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (c *MollieCapture) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (c *MollieCapture) Filterable() bool {
	return true
}

func capturesColMap() map[string]string {
	return map[string]string{
		"RESOURCE":          "the resource name",
		"ID":                "the resource id",
		"MODE":              "the mode used to create this capture",
		"AMOUNT":            "the amount captured",
		"SETTLEMENT_AMOUNT": "the amount that will be settled to your account",
		"PAYMENT_ID":        "the unique identifier of the payment",
		"SHIPMENT_ID":       "the unique identifier of the shipment",
		"SETTLEMENT_ID":     "the unique identifier of the settlement",
		"CREATED_AT":        "the capture’s date and time of creation",
	}
}

func capturesCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"AMOUNT",
		"SETTLEMENT_AMOUNT",
		"PAYMENT_ID",
		"SHIPMENT_ID",
		"SETTLEMENT_ID",
		"CREATED_AT",
	}
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
