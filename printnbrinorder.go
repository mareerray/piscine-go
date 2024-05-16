package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	if n < 0 {
		z01.PrintRune(0)
	}
	runes := ToRunes(n)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	for i := range runes {
		z01.PrintRune(runes[i])
	}
}

func ToRunes(n int) []rune {
	runes := make([]rune, 0)
	for n != 0 {
		num := n % 10
		n = n / 10
		runes = append(runes, rune(num+'0'))
	}
	return runes
}
