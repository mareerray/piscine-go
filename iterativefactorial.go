package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	var result int
	result = 1

	for i := 1; i <= nb; i++ {
		result = result * i
	}

	if nb >= 21 {
		return 0
	}

	return result
}
