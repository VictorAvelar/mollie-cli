package displayers

import (
	"strings"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
)

const (
	displayDateFormat = "02-01-2006"
	noDateContent     = "----------"
	noAppFee          = "none"
)

func fallbackSafeDate(t *time.Time) string {
	if t == nil {
		return noDateContent
	}

	return t.Format(displayDateFormat)
}

func fallbackSafeAmount(a *mollie.Amount) string {
	if a == nil {
		return "--- ---"
	}

	return stringCombinator(" ", a.Value, a.Currency)
}

func fallbackSafePaymentMethod(pm mollie.PaymentMethod) string {
	if pm == mollie.PaymentMethod("") {
		return "none"
	}

	return string(pm)
}

func fallbackSafeAppFee(af *mollie.ApplicationFee) string {
	if af == nil {
		return noAppFee
	}

	return fallbackSafeAmount(af.Amount)
}

func stringCombinator(s string, parts ...string) string {
	for i, v := range parts {
		if v == "" {
			parts[i] = "-"
		}
	}
	return strings.Join(parts, s)
}
