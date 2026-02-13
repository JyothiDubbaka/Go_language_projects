package main

import "fmt"

func ClearScreen() {
	// ANSI escape code to clear screen and move cursor to top-left
	//033 ESC charater
	//[2J Clear entire screen
	// [H Move cursor to home position
	fmt.Print("\033[H\033[2J")
}

//Introducing colours
const (
	Reset  = "\033[0m"
	Bold   = "\033[1m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
)