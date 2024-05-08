package piscine

import "github.com/01-edu/z01"

func PrintComb() {

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

	for a := 0; a < 10; a++ {
		for b := 0; b < 10; b++ {
			for c := 0; c < 10; c++ {
				if a < b && b < c {
					z01.PrintRune(RuneMap[a])
					z01.PrintRune(RuneMap[b])
					z01.PrintRune(RuneMap[c])
					if a == 7 && b == 8 && c == 9 {
						z01.PrintRune('\n')
						continue
					}
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}

			}
		}
	}
}
