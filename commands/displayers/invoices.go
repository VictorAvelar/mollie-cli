package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieInvoice wrapper for displaying.
type MollieInvoice struct {
	*mollie.Invoice
}

// KV is a displayable group of key value.
func (mi *MollieInvoice) KV() []map[string]interface{} {
	var out []map[string]interface{}
	x := buildXInvoice(mi.Invoice)
	out = append(out, x)
	return out
}

// MollieInvoiceList wrapper for displaying.
type MollieInvoiceList struct {
	*mollie.InvoiceList
}

// KV is a displayable group of key value.
func (mil *MollieInvoiceList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, i := range mil.Embedded.Invoices {
		x := buildXInvoice(&i)

		out = append(out, x)
	}

	return out
}

func buildXInvoice(i *mollie.Invoice) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":     i.Reference,
		"ID":           i.ID,
		"REFERENCE":    i.Reference,
		"VAT_NUMBER":   i.VatNumber,
		"STATUS":       i.Status,
		"ISSUED_AT":    i.IssuedAt,
		"PAID_AT":      i.PaidAt,
		"DUE_AT":       i.DueAt,
		"NET_AMOUNT":   fallbackSafeAmount(i.NetAmount),
		"VAT_AMOUNT":   fallbackSafeAmount(i.VatAmount),
		"GROSS_AMOUNT": fallbackSafeAmount(i.GrossAmount),
	}
}
