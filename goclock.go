package main

import (
	"flag"
	"fmt"
	"strings"
	"time"

	terminal "github.com/wayneashleyberry/terminal-dimensions"
)

const clearCode string = "\033[H\033[2J"

var colorCode = map[string]int{
	"red":     31,
	"green":   32,
	"yellow":  33,
	"blue":    34,
	"magenta": 35,
	"cyan":    36,
	"white":   37,
}

func getColorCode(color string, bg bool) string {
	if bg {
		return "\u001b[" + fmt.Sprint(colorCode[color]+10) + "m"
	} else {
		return "\u001b[" + fmt.Sprint(colorCode[color]) + "m"
	}
}

var full = [6]bool{true, true, true, true, true, true}
var empty = [6]bool{false, false, false, false, false, false}
var left = [6]bool{true, true, false, false, false, false}
var right = [6]bool{false, false, false, false, true, true}
var outer = [6]bool{true, true, false, false, true, true}
var inner = [6]bool{false, false, true, true, false, false}

var characters = map[rune][6][6]bool{
	'0': {full, outer, outer, outer, full},
	'1': {inner, inner, inner, inner, inner},
	'2': {full, right, full, left, full},
	'3': {full, right, full, right, full},
	'4': {outer, outer, full, right, right},
	'5': {full, left, full, right, full},
	'6': {full, left, full, outer, full},
	'7': {full, right, right, right, right},
	'8': {full, outer, full, outer, full},
	'9': {full, outer, full, right, right},
	':': {empty, inner, empty, inner, empty}}

func main() {
	var colorFlag = flag.String("color", "white", "The color to use.\nOptions: 'red', 'green', 'yellow', 'blue', 'magenta', 'cyan', 'white'")
	var centerFlag = flag.Bool("center", false, "Center the clock on the terminal.")
	flag.Parse()
	var bgColor = getColorCode(*colorFlag, true)
	var color = getColorCode(*colorFlag, false)

	for {
		var xOffset = ""
		var yOffset = ""
		if *centerFlag {
			var x, _ = terminal.Width()
			var y, _ = terminal.Height()
			xOffset = strings.Repeat(" ", int(x/2-19))
			yOffset = strings.Repeat("\n", int(y/2-2))
		}

		var currentTime = time.Now()
		var formattedTime = currentTime.Format("15:04")           //hh:mm
		var formattedDate = currentTime.Format("_2 January 2006") //day month year

		//var formattedTime = "0123456789:" //test font

		var output = yOffset

		for i := 0; i < 5; i++ {
			output += xOffset
			for _, c := range formattedTime {
				for ii := 0; ii < 6; ii++ {
					if characters[c][i][ii] {
						output += bgColor + " \u001b[0m"
					} else {
						output += " "
					}
				}
				output += "  "
			}
			output += "\n"
		}
		output += "\n" + color + xOffset + formattedDate

		fmt.Print(
			clearCode,
			output,
		)
		time.Sleep(1 * time.Second)
	}
}
