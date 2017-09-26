// Basic Go stuff,  myfernandez@gmail.com
/*
LogPrint the season a month belongs to. The month is provided by user or assigned to a variable.  Implement the solution using :
ifthenelse and switchcase.
Winter = Nov, Dec, Jan, Feb.
Spring   = March, April, May.
Summer = Jun, July, Aug,
Fall = Sept, Oct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	m2s "github.com/myfernandez/month2season"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Month: ")
	scanner.Scan()
	month := scanner.Text()
	season, err := m2s.ReturnSeason(month)
	if err != nil {
		log.Fatal("The season is unknown")
	}
	fmt.Printf("The season is: %s\n", season)
}
