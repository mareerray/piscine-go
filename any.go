package piscine

func Any(f func(string) bool, a []string) bool {
	// passing item from the slice through the fucntion IsNumeric

	for _, value := range a {
		if f(value) {
			return true
		}
	}
	return false
}
