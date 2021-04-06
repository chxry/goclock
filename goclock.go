package main

import (
	"flag"
	"fmt"
	"time"
)

const clearCode string = "\033[H\033[2J"

var colorCode = map[string]string{
	"red":     "\u001b[41m",
	"green":   "\u001b[42m",
	"yellow":  "\u001b[43m",
	"blue":    "\u001b[44m",
	"magenta": "\u001b[45m",
	"cyan":    "\u001b[46m",
	"white":   "\u001b[47m",
}

var full = [6]bool{true, true, true, true, true, true}
var empty = [6]bool{false, false, false, false, false, false}
var left = [6]bool{true, true, false, false, false, false}
var right = [6]bool{false, false, false, false, true, true}
var outer = [6]bool{true, true, false, false, true, true}
var inner = [6]bool{false, false, true, true, false, false}

var characters = map[rune][5][6]bool{
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
	var color = flag.String("color", "white", "The color to use.\nOptions: 'red', 'green', 'yellow', 'blue', 'magenta', 'cyan', 'white'")
	flag.Parse()

	for {
		var currentTime = time.Now().Format("15:04") //hh:mm
		//var currentTime = "0123456789:" //test font
		var output = ""
		for i := 0; i < 5; i++ {
			for _, c := range currentTime {
				for ii := 0; ii < 6; ii++ {
					if characters[c][i][ii] {
						output += colorCode[*color] + " \u001b[0m"
					} else {
						output += " "
					}
				}
				output += "  "
			}
			output += "\n"
		}
		fmt.Print(
			clearCode,
			output,
		)
		time.Sleep(1 * time.Second)
	}
}
