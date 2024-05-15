package piscine

func FirstRune(s string) rune {
	for _, one := range s {
		return one
	}
	return 0
}
