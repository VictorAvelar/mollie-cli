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
