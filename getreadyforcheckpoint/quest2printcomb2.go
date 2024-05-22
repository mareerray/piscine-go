package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	RuneMap := map[int]rune{
		0: '0',
		1: '1',
		2: '2',
		3: '3',
		4: '4',
		5: '5',
		6: '6',
		7: '7',
		8: '8',
		9: '9',
	}
	for a := 0; a <= 99; a++ { //iterates through all possible first digits of two-digit numbers (a), starting from 0 to 99.
		for b := a + 1; b <= 99; b++ { //iterates through all possible second digits of two-digit numbers (b), starting from a+1 to ensure each pair consists of two different digits.
			z01.PrintRune(RuneMap[a/10])
			z01.PrintRune(RuneMap[a%10])
			z01.PrintRune(' ')
			z01.PrintRune(RuneMap[b/10])
			z01.PrintRune(RuneMap[b%10])
			// z01.PrintRune(',')
			// z01.PrintRune(' ')

			if b == 99 && a == 98 {
				z01.PrintRune('\n')
				break
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}

	}

}
