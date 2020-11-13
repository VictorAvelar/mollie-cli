package displayers

import (
	"testing"
	"time"

	"github.com/VictorAvelar/mollie-api-go/v2/mollie"
	"github.com/stretchr/testify/assert"
)

func TestMollieProfile_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	disp := MollieProfile{
		Profile: &mollie.Profile{
			ID:        "pr_test",
			Name:      "testing profile",
			Website:   "https://example.com",
			Phone:     "+0000000000000",
			Status:    mollie.StatusVerified,
			Mode:      mollie.TestMode,
			CreatedAt: &n,
		},
	}

	out := expectProfileSlice(disp.Profile)
	assert.Len(t, out, 1)
	assert.Equal(t, out, disp.KV())
}

func TestMollieProfileList_KV(t *testing.T) {
	n, err := time.Parse("02-01-2006", "01-11-2020")
	if err != nil {
		t.Error(err)
	}
	var ps []*mollie.Profile
	{
		ps = append(ps, &mollie.Profile{
			ID:        "pr_test",
			Name:      "testing profile",
			Website:   "https://example.com",
			Phone:     "+0000000000000",
			Status:    mollie.StatusVerified,
			Mode:      mollie.TestMode,
			CreatedAt: &n,
		},
			&mollie.Profile{
				ID:        "pr_test_2",
				Name:      "testing profile 2",
				Website:   "https://example.com/2",
				Phone:     "+0000000000000",
				Status:    mollie.StatusUnverified,
				Mode:      mollie.LiveMode,
				CreatedAt: &n,
			},
		)
	}

	disp := MollieProfileList{
		ProfileList: &mollie.ProfileList{
			Count: 2,
			Embedded: struct {
				Profiles []*mollie.Profile "json:\"profiles,omitempty\""
			}{
				Profiles: ps,
			},
			Links: mollie.PaginationLinks{
				Documentation: mollie.URL{Href: "https://example.com", Type: "text/html"},
				Self:          mollie.URL{Href: "https://example.com", Type: "text/html"},
			},
		},
	}

	out := expectProfileSlice(disp.Embedded.Profiles...)

	assert.Len(t, out, 2)
	assert.Equal(t, disp.Count, 2)
	assert.Equal(t, out, disp.KV())
}

func expectProfileSlice(pfs ...*mollie.Profile) (out []map[string]interface{}) {
	for _, r := range pfs {
		x := map[string]interface{}{
			"ID":      r.ID,
			"Name":    r.Name,
			"Website": r.Website,
			"Phone":   r.Phone,
			"Status":  r.Status,
			"Mode":    r.Mode,
			"Since":   r.CreatedAt.Format("02-01-2006"),
		}

		out = append(out, x)
	}

	return
}
