package displayers

import (
	"strings"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

func stringCombinator(s string, parts ...string) string {
	for i, v := range parts {
		if v == "" {
			parts[i] = "-"
		}
	}
	return strings.Join(parts, s)
}

func getSafeExpiration(p mollie.Payment) string {
	if p.ExpiresAt == nil {
		return "----------"
	}

	return p.ExpiresAt.Format("02-01-2006")
}

func getSafePaymentMethod(p mollie.Payment) string {
	if p.Method == mollie.PaymentMethod("") {
		return "none"
	}
	return string(p.Method)
}
