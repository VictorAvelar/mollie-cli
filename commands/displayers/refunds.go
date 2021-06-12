package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieRefundList wrapper for displaying.
type MollieRefundList struct {
	*mollie.RefundList
}

// KV is a displayable group of key value.
func (mrl *MollieRefundList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, r := range mrl.Embedded.Refunds {
		x := buildXRefund(r)

		out = append(out, x)
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (mrl *MollieRefundList) Cols() []string {
	return refundsCols()
}

// ColMap returns a list of columns and its description.
func (mrl *MollieRefundList) ColMap() map[string]string {
	return refundsColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mrl *MollieRefundList) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mrl *MollieRefundList) Filterable() bool {
	return true
}

// MollieRefund wrapper for displaying.
type MollieRefund struct {
	*mollie.Refund
}

// KV is a displayable group of key value.
func (mr *MollieRefund) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXRefund(mr.Refund)

	out = append(out, x)

	return out
}

// Cols returns an array of columns available for displaying.
func (mr *MollieRefund) Cols() []string {
	return refundsCols()
}

// ColMap returns a list of columns and its description.
func (mr *MollieRefund) ColMap() map[string]string {
	return refundsColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mr *MollieRefund) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mr *MollieRefund) Filterable() bool {
	return true
}

func refundsCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"AMOUNT",
		"SETTLEMENT_ID",
		"SETTLEMENT_AMOUNT",
		"DESCRIPTION",
		"METADATA",
		"STATUS",
		"PAYMENT_ID",
		"ORDER_ID",
		"CREATED_AT",
	}
}

func refundsColMap() map[string]string {
	return map[string]string{
		"RESOURCE":          "the resource name",
		"ID":                "the resource id",
		"AMOUNT":            "the amount refunded to your customer",
		"SETTLEMENT_ID":     "the identifier referring to the settlement this payment was settled with",
		"SETTLEMENT_AMOUNT": "the amount that will be deducted from your account balance",
		"DESCRIPTION":       "the description of the refund that may be shown to your customer,",
		"METADATA":          "metadata you provided upon refund creation",
		"STATUS":            "the refund carries a status field",
		"PAYMENT_ID":        "the unique identifier of the payment this refund was created for",
		"ORDER_ID":          "the unique identifier of the order this refund was created for",
		"CREATED_AT":        "the date and time the refund was issued",
	}
}

func buildXRefund(r *mollie.Refund) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":          r.Resource,
		"ID":                r.ID,
		"AMOUNT":            fallbackSafeAmount(r.Amount),
		"SETTLEMENT_ID":     r.SettlementID,
		"SETTLEMENT_AMOUNT": fallbackSafeAmount(r.SettlementAmount),
		"DESCRIPTION":       r.Description,
		"METADATA":          r.Metadata,
		"STATUS":            r.Status,
		"PAYMENT_ID":        r.PaymentID,
		"ORDER_ID":          r.OrderID,
		"CREATED_AT":        fallbackSafeDate(r.CreatedAt),
	}
}
