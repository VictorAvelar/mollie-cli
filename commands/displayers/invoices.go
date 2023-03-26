package displayers

import "github.com/VictorAvelar/mollie-api-go/v3/mollie"

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

// Cols returns an array of columns available for displaying.
func (mi *MollieInvoice) Cols() []string {
	return invoicesCols()
}

// ColMap returns a list of columns and its description.
func (mi *MollieInvoice) ColMap() map[string]string {
	return invoicesColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mi *MollieInvoice) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mi *MollieInvoice) Filterable() bool {
	return true
}

// MollieInvoiceList wrapper for displaying.
type MollieInvoiceList struct {
	*mollie.InvoicesList
}

// KV is a displayable group of key value.
func (mil *MollieInvoiceList) KV() []map[string]interface{} {
	out := outPrealloc(len(mil.Embedded.Invoices))

	for i := range mil.Embedded.Invoices {
		out = append(out, buildXInvoice(&mil.Embedded.Invoices[i]))
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (mil *MollieInvoiceList) Cols() []string {
	return invoicesCols()
}

// ColMap returns a list of columns and its description.
func (mil *MollieInvoiceList) ColMap() map[string]string {
	return invoicesColMap()
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mil *MollieInvoiceList) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mil *MollieInvoiceList) Filterable() bool {
	return true
}

func invoicesCols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"REFERENCE",
		"VAT_NUMBER",
		"STATUS",
		"ISSUED_AT",
		"PAID_AT",
		"DUE_AT",
		"NET_AMOUNT",
		"VAT_AMOUNT",
		"GROSS_AMOUNT",
	}
}

func invoicesColMap() map[string]string {
	return map[string]string{
		"RESOURCE":     "the resource name",
		"ID":           "the resource id",
		"REFERENCE":    "a specific invoice number / reference",
		"VAT_NUMBER":   "invoices from a specific year",
		"STATUS":       "status of the invoice",
		"ISSUED_AT":    "the invoice date",
		"PAID_AT":      "the date on which the invoice was paid",
		"DUE_AT":       "the date on which the invoice is due",
		"NET_AMOUNT":   "total amount of the invoice excluding VAT",
		"VAT_AMOUNT":   "VAT amount of the invoice",
		"GROSS_AMOUNT": "total amount of the invoice including VAT",
	}
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
