package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieChargeback wrapper for displaying
type MollieChargeback struct {
	*mollie.Chargeback
}

// KV is a displayable group of key value
func (cb *MollieChargeback) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXChargeback(cb.Chargeback)

	out = append(out, x)

	return out
}

// MollieChargebackList wrapper for displaying
type MollieChargebackList struct {
	*mollie.ChargebackList
}

// KV is a displayable group of key value
func (lp *MollieChargebackList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, p := range lp.Embedded.Chargebacks {
		x := buildXChargeback(&p)

		out = append(out, x)
	}

	return out
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
