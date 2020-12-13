package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieProfileList wrapper for displaying.
type MollieProfileList struct {
	*mollie.ProfileList
}

// KV is a displayable group of key value
func (mpl *MollieProfileList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, r := range mpl.Embedded.Profiles {
		x := buildXProfile(r)

		out = append(out, x)
	}

	return out
}

// MollieProfile wrapper for displaying
type MollieProfile struct {
	*mollie.Profile
}

// KV is a displayable group of key value
func (mp *MollieProfile) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := buildXProfile(mp.Profile)

	out = append(out, x)

	return out
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
