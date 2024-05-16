package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// var arg("./")# command-line-arguments
	//   ./main.go:9:8: syntax error: unexpected ., expected type
	// appName := os.Arg[0]
	// z01.PrintRune("appName")
	// 	cannot use "appName" (untyped string constant) as rune value in argument to z01.PrintRune
	// mayuree.reunsati@Y6WJLXCPYF printprogramname %

	// get the program name using os.Args
	appName := os.Args[0]
	// find the last seperator in the path program ("/" seperator main.go)
	seperator := '/'

	lastSep := -1
	for i, c := range appName {
		if c == seperator {
			lastSep = i
		}
	}
	// extract the base name from path or example safari instead of safari.app remove extension
	simple := appName
	if lastSep != -1 {
		simple = appName[lastSep+1:]
	}
	// print each character using z01.PrintRune
	for _, r := range simple {
		z01.PrintRune(r)
	}
	// now this pring a new line

	z01.PrintRune('\n')
}
