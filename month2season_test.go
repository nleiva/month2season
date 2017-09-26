package month2season_test

import (
	"testing"

	m2s "github.com/myfernandez/month2season"
)

// ReturnSeason returns the matching season
func TestReturnSeason(t *testing.T) {
	tt := []struct {
		name   string
		month  []string
		season string
		err    string
	}{
		{name: "January", month: []string{"Jan", "jan", "january"}, season: "Winter"},
		{name: "February", month: []string{"Feb", "feb", "february"}, season: "Winter"},
		{name: "March", month: []string{"Mar", "mar", "march"}, season: "Spring"},
		{name: "April", month: []string{"Apr", "april", "april"}, season: "Spring"},
		{name: "May", month: []string{"May", "may"}, season: "Spring"},
		{name: "June", month: []string{"Jun", "jun", "june"}, season: "Summer"},
		{name: "July", month: []string{"Jul", "jul", "july"}, season: "Summer"},
		{name: "August", month: []string{"Aug", "aug", "august"}, season: "Summer"},
		{name: "September", month: []string{"Sep", "sep", "september"}, season: "Fall"},
		{name: "October", month: []string{"Oct", "oct", "october"}, season: "Fall"},
		{name: "November", month: []string{"Nov", "nov", "november"}, season: "Fall"},
		{name: "December", month: []string{"Dec", "dec", "december"}, season: "Winter"},
	}
	for _, tc := range tt {
		for _, m := range tc.month {
			_, err := m2s.ReturnSeason(m)
			if err != nil {
				if err.Error() != tc.err {
					t.Fatalf("could not match %s to %s", m, tc.season)
				}
			}
		}
	}
}
