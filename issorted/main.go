package main

import (
	"fmt"

	"piscine"
)

// Define the comparison function f
func f(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func main() {
	a1 := []int{0, 1, 2, 3, 4, 5}
	a2 := []int{0, 2, 1, 3}
	a3 := []int{9, 7, 5, 5, -3}

	result1 := piscine.IsSorted(f, a1)
	result2 := piscine.IsSorted(f, a2)
	result3 := piscine.IsSorted(f, a3)

	fmt.Println(result1) // Should print: true
	fmt.Println(result2) // Should print: false
	fmt.Println(result3)
}
