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
		x := map[string]interface{}{
			"ID":      r.ID,
			"Name":    r.Name,
			"Website": r.Website,
			"Phone":   r.Phone,
			"Status":  r.Status,
			"Mode":    r.Mode,
			"Since":   r.CreatedAt.Format("01-02-2006"),
		}

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

	x := map[string]interface{}{
		"ID":      mp.Profile.ID,
		"Name":    mp.Profile.Name,
		"Website": mp.Profile.Website,
		"Phone":   mp.Profile.Phone,
		"Status":  mp.Profile.Status,
		"Mode":    mp.Profile.Mode,
		"Since":   mp.CreatedAt.Format("01-02-2006"),
	}

	out = append(out, x)

	return out
}
