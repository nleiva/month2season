/*
Basic Go stuff,  myfernandez@gmail.com

LogPrint the season a month belongs to. The month is provided by user or assigned to a variable.

Implements the solution using: ifthenelse and switchcase.
 Winter = Nov, Dec, Jan, Feb.
 Spring   = March, April, May.
 Summer = Jun, July, Aug,
 Fall = Sept, Oct.
*/

package month2season

import (
	"fmt"
	"strings"
)

// ReturnSeason returns the matching season
func ReturnSeason(m string) (string, error) {
	m = strings.ToLower(m)
	switch {
	case strings.Contains(m, "dec"), strings.Contains(m, "jan"), strings.Contains(m, "feb"):
		return "Winter", nil
	case strings.Contains(m, "mar"), strings.Contains(m, "apr"), strings.Contains(m, "may"):
		return "Spring", nil
	case strings.Contains(m, "jun"), strings.Contains(m, "jul"), strings.Contains(m, "aug"):
		return "Summer", nil
	case strings.Contains(m, "sep"), strings.Contains(m, "oct"), strings.Contains(m, "nov"):
		return "Fall", nil
	default:
		return "", fmt.Errorf("month %s is not valid", m)
	}
}
