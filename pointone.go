package piscine

import "fmt"

func PointOne(n *int) {
	fmt.Println("value in n: %v ", *n)
	*n = 1 // * is used to tell computer to handle the value,
	// assign 1 to value of a
	fmt.Println("address in n: %v ", n)
	fmt.Println("new value in n: %v ", *n)
}
