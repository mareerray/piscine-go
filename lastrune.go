package piscine

func LastRune(s string) rune {
	last := []rune(s)   // convert from a string to a slice of runes
	if len(last) == 0 { // returns the number of elements in the runes slice.
		return 0 // If the slice is empty (i.e., the length is 0), the function returns 0.
	}
	return last[len(last)-1] // accesses the last element of the slice, and return it
}
