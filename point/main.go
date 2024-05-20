package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(points *point) {
	points.x = 42
	points.y = 21
}

func printNumber(s int) {
	if s == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	if s < 0 {
		z01.PrintRune('-')
		s = -s
	}
	for s > 0 {
		digits = append([]rune{rune('0' + s%10)}, digits...)
		s /= 10
	}
	for _, digits := range digits {
		z01.PrintRune(digits)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune(rune('x'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	printNumber(points.x)
	z01.PrintRune(rune(','))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('y'))
	z01.PrintRune(rune(' '))
	z01.PrintRune(rune('='))
	z01.PrintRune(rune(' '))
	printNumber(points.y)
	z01.PrintRune('\n')
}
