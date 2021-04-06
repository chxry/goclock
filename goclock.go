package main

import (
	"fmt"
	"time"
)

const clearCode string = "\033[H\033[2J"

var characters = map[rune][5]string{
	'0': {"██████", "█    █", "█    █", "█    █", "██████"},
	'1': {"  ██  ", "  ██  ", "  ██  ", "  ██  ", "  ██  "},
	'2': {"██████", "     █", "██████", "█     ", "██████"},
	'3': {"██████", "     █", "██████", "     █", "██████"},
	'4': {"█    █", "█    █", "██████", "     █", "     █"},
	'5': {"██████", "█     ", "██████", "     █", "██████"},
	'6': {"██████", "█     ", "██████", "█    █", "██████"},
	'7': {"██████", "     █", "     █", "     █", "     █"},
	'8': {"██████", "█    █", "██████", "█    █", "██████"},
	'9': {"██████", "█    █", "██████", "     █", "     █"},
	':': {"      ", "  ██  ", "      ", "  ██  ", "      "}}

func main() {
	for {
		//var x, _ = terminal.Width()
		var currentTime = time.Now().Format("15:04") //hh:mm
		var output = ""

		for i := 0; i < 5; i++ {
			for _, c := range currentTime {
				output += characters[c][i]
				output += "  "
			}
			output += "\n"
		}

		fmt.Print(
			clearCode,
			//strings.Repeat("#", int(x/2)),
			output,
		)
		time.Sleep(1 * time.Second)
	}
}
