package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argument := os.Args[1:]
	if len(argument) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(argument) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	fileName := argument[0]
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Failed to read file", err)
	}
	// Convert the byte slice to a string and print it
	fmt.Print(string(content))
}
