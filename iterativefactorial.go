package piscine

func IterativeFactorial(nb int) int {

	if nb < 0 {
		return 0
	}
	var result int
	result = 0

	for i := 0; i < nb+1; i++ {
		result = result + i
	}
	return result

}
