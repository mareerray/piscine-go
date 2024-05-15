package piscine

func FirstRune(s string) rune { // function FirstRune iterates over the string s using a range-based for loop.

	for _, one := range s { // The loop returns the first rune (one) it encounters.

		return one // The loop exits immediately after returning the first rune.
	}
	return 0 // This ensures the function returns a valid rune
}

// Goal of the task: function that returns the first rune of a string.