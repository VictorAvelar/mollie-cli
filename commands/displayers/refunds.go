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
		x := map[string]interface{}{
			"ID":          r.ID,
			"Payment":     r.PaymentID,
			"Order":       r.OrderID,
			"Settlement":  r.SettlementID,
			"Amount":      stringCombinator(" ", r.Amount.Value, r.Amount.Currency),
			"Status":      r.Status,
			"Description": r.Description,
			"Created at":  r.CreatedAt.Format("02-01-2006"),
		}

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

	x := map[string]interface{}{
		"ID":          mr.ID,
		"Payment":     mr.PaymentID,
		"Order":       mr.OrderID,
		"Settlement":  mr.SettlementID,
		"Amount":      stringCombinator(" ", mr.Amount.Value, mr.Amount.Currency),
		"Status":      mr.Status,
		"Description": mr.Description,
		"Created at":  mr.CreatedAt.Format("02-01-2006"),
	}

	out = append(out, x)

	return out
}
