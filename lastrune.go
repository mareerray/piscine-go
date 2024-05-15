package piscine

func LastRune(s string) rune {
	last := []rune(s)   // convert from a string to a slice of runes
	if len(last) == 0 { // returns the number of elements in the runes slice.
		return 0 // If the slice is empty (i.e., the length is 0), the function returns 0.
	}
	return last[len(last)-1] // accesses the last element of the slice, and return it
}

/*Goal of the task:function that returns the last rune of a string.
task1 : convert from a string to a slice of runes to ensure that
each element in the slice represents one Unicode character.
task2 : if statement checks if the length of the runes slice is zero.
task3 : runes[len(runes)-1] accesses the last element of the slice,
which is the last rune (character) of the input string s.
*/
