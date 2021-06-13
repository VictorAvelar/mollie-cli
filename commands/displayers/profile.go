package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieProfileList wrapper for displaying.
type MollieProfileList struct {
	*mollie.ProfileList
}

// KV is a displayable group of key value.
func (mpl *MollieProfileList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, r := range mpl.Embedded.Profiles {
		x := buildXProfile(r)

		out = append(out, x)
	}

	return out
}

// MollieProfile wrapper for displaying.
type MollieProfile struct {
	*mollie.Profile
}

// KV is a displayable group of key value.
func (mp *MollieProfile) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXProfile(mp.Profile)

	out = append(out, x)

	return out
}

// Cols returns an array of columns available for displaying.
func (mp *MollieProfile) Cols() []string {
	return []string{
		"RESOURCE",
		"ID",
		"MODE",
		"NAME",
		"WEBSITE",
		"EMAIL",
		"PHONE",
		"CATEGORY_CODE",
		"STATUS",
		"REVIEW",
		"CREATED_AT",
	}
}

// ColMap returns a list of columns and its description.
func (mp *MollieProfile) ColMap() map[string]string {
	return map[string]string{
		"RESOURCE":      "the resource name",
		"ID":            "the resource id",
		"MODE":          "the profile mode (live/test)",
		"NAME":          "the profile name",
		"WEBSITE":       "the profile website",
		"EMAIL":         "the profile registered email",
		"PHONE":         "the profile registered phone number",
		"CATEGORY_CODE": "the profile category code (see mollie categories)",
		"STATUS":        "the profile status",
		"REVIEW":        "the profile review status",
		"CREATED_AT":    "the profile creation date",
	}
}

// NoHeaders returns a boolean indicating if headers should be displayed
// or not to the provided output.
func (mp *MollieProfile) NoHeaders() bool {
	return false
}

// Filterable indicates if the displayable output can be filtered
// using the fields flag.
func (mp *MollieProfile) Filterable() bool {
	return true
}

func buildXProfile(p *mollie.Profile) map[string]interface{} {
	return map[string]interface{}{
		"RESOURCE":      p.Resource,
		"ID":            p.ID,
		"MODE":          fallbackSafeMode(p.Mode),
		"NAME":          p.Name,
		"WEBSITE":       p.Website,
		"EMAIL":         p.Email,
		"PHONE":         p.Phone,
		"CATEGORY_CODE": p.CategoryCode,
		"STATUS":        p.Status,
		"REVIEW":        p.Review.Status,
		"CREATED_AT":    fallbackSafeDate(p.CreatedAt),
	}
}
