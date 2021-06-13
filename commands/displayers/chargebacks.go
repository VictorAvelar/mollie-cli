package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieChargeback wrapper for displaying.
type MollieChargeback struct {
	*mollie.Chargeback
}

// KV is a displayable group of key value.
func (cb *MollieChargeback) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXChargeback(cb.Chargeback)

	out = append(out, x)

	return out
}

// Cols returns an array of columns available for displaying.
func (cb *MollieChargeback) Cols() []string {
	return chargebackCols()
}

// ColMap returns a list of columns and its description.
func (cb *MollieChargeback) ColMap() map[string]string {
	return chargebackColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed.
// or not to the provided output.
func (cb *MollieChargeback) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (cb *MollieChargeback) Filterable() bool {
	return true
}

// MollieChargebackList wrapper for displaying.
type MollieChargebackList struct {
	*mollie.ChargebackList
}

// KV is a displayable group of key value.
func (lp *MollieChargebackList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for i := range lp.Embedded.Chargebacks {
		out = append(out, buildXChargeback(&lp.Embedded.Chargebacks[i]))
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (lp *MollieChargebackList) Cols() []string {
	return chargebackCols()
}

// ColMap returns a list of columns and its description.
func (lp *MollieChargebackList) ColMap() map[string]string {
	return chargebackColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed.
// or not to the provided output.
func (lp *MollieChargebackList) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (lp *MollieChargebackList) Filterable() bool {
	return true
}

func chargebackCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"AMOUNT",
		"SETTLEMENT_AMOUNT",
		"CREATED_AT",
		"REVERSED_AT",
		"PAYMENT_ID",
	}
}

func chargebackColMap() map[string]string {
	return map[string]string{
		"RESOURCE":          "the resource name",
		"ID":                "the resource id",
		"AMOUNT":            "the amount charged back by the consumer",
		"SETTLEMENT_AMOUNT": "the amount that will be deducted from your account",
		"CREATED_AT":        "the date and time the chargeback was issued",
		"REVERSED_AT":       "the date and time the chargeback was reversed if applicable",
		"PAYMENT_ID":        "the unique identifier of the payment this chargeback was issued for",
	}
}

func buildXChargeback(cb *mollie.Chargeback) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":          cb.Resource,
		"ID":                cb.ID,
		"AMOUNT":            fallbackSafeAmount(cb.Amount),
		"SETTLEMENT_AMOUNT": fallbackSafeAmount(cb.SettlementAmount),
		"CREATED_AT":        fallbackSafeDate(cb.CreatedAt),
		"REVERSED_AT":       fallbackSafeDate(cb.ReversedAt),
		"PAYMENT_ID":        cb.PaymentID,
	}
}
