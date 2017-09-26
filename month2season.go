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
import "fmt"
import "bufio"
import "os"
import "strings"

func ReadTextStdin() string {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Month (eg Jan, Feb, Mar, etc:) ")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    fmt.Println("Your month is: ", text)
    //text := bufio.NewScanner(os.Stdin)
    strings.TrimSpace(text)
    return text
}

func ReturnSeason(month string) string {
	if strings.ToUpper(month) == strings.ToUpper("Nov") {
	   return "Winter"
	} else {
	   switch strings.ToUpper(month) {
		case strings.ToUpper("Dec"), strings.ToUpper("Jan"), strings.ToUpper("Feb"):
			return "Winter"  
		case strings.ToUpper("Mar"), strings.ToUpper("Apr"), strings.ToUpper("May"):
			return "Spring"
		case strings.ToUpper("Jun"), strings.ToUpper("Jul"), strings.ToUpper("Aug"):
			return "Summer"
		default:
			return "Fall"
		}
	   }
}


func main() {

// 
	month:= ReadTextStdin()
	fmt.Println("The season is: ", ReturnSeason(month))
}
