package piscine

func Map(f func(int) bool, a []int) []bool {
	// Initialize a slice of booleans with the same length as `a`
	b := make([]bool, len(a))

	// Apply `f` to each element of `a` and store the result in `b`
	for i := range a {
		b[i] = f(a[i])
	}

	// Return the slice of booleans
	return b
}
