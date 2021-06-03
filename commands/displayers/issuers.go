package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MolliePaymentMethodIssuer wrapper for displaying.
type MolliePaymentMethodIssuer struct {
	*mollie.PaymentMethodIssuer
}

// KV is a displayable group of key value.
func (mpi *MolliePaymentMethodIssuer) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXIssuer(mpi.PaymentMethodIssuer)

	out = append(out, x)

	return out
}

// Cols returns an array of columns available for displaying.
func (mpi *MolliePaymentMethodIssuer) Cols() []string {
	return []string{
		"ID",
		"NAME",
		"IMAGE",
	}
}

// ColMap returns a list of columns and its description.
func (mpi *MolliePaymentMethodIssuer) ColMap() map[string]string {
	return map[string]string{
		"ID":    "the issuer id",
		"NAME":  "the issuer name",
		"IMAGE": "the issuer logo/image",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mpi *MolliePaymentMethodIssuer) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mpi *MolliePaymentMethodIssuer) Filterable() bool {
	return true
}

// MollieListPaymentMethodsIssuers wrapper for displaying.
type MollieListPaymentMethodsIssuers struct {
	Issuers []*mollie.PaymentMethodIssuer
}

// KV is a displayable group of key value.
func (lpmi *MollieListPaymentMethodsIssuers) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, pm := range lpmi.Issuers {
		x := buildXIssuer(pm)

		out = append(out, x)
	}

	return out
}

// Cols returns an array of columns available for displaying.
func (lpmi *MollieListPaymentMethodsIssuers) Cols() []string {
	return []string{
		"ID",
		"NAME",
		"IMAGE",
	}
}

// ColMap returns a list of columns and its description.
func (lpmi *MollieListPaymentMethodsIssuers) ColMap() map[string]string {
	return map[string]string{
		"ID":    "the issuer id",
		"NAME":  "the issuer name",
		"IMAGE": "the issuer logo/image",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (lpmi *MollieListPaymentMethodsIssuers) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (lpmi *MollieListPaymentMethodsIssuers) Filterable() bool {
	return true
}

func buildXIssuer(mi *mollie.PaymentMethodIssuer) map[string]interface{} {
	return map[string]interface{}{
		"ID":    mi.ID,
		"NAME":  mi.Name,
		"IMAGE": mi.Image.Svg,
	}
}
