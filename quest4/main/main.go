package main

import (
	"fmt"
	"piscine"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
	// or we can do like this
	// result := ConcatParams(arg)
	// fmt.Println(result)
}
