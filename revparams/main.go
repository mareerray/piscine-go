package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// Reverse the order of the arguments
	for i := len(args) - 1; i >= 0; i-- {
		arg := args[i]

		// Print each character of the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
