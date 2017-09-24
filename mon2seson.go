// Basic Go Stuff, myfernandez@gmail.com
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
    fmt.Println("Your month is: ", text)
    //text := bufio.NewScanner(os.Stdin)
    strings.TrimSpace(text)
    return text
}

func ReturnSeason(month string) string {
	if month == "Nov" {
	   return "Winter"
	} else {
	   switch month {
		case "Dec", "Jan", "Feb":
		return "Winter"  
		case "Mar", "Apr", "May":
		return "Spring"
		case "Jun","Jul", "Aug":
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
