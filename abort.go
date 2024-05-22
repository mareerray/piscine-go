package piscine

func Abort(a, b, c, d, e int) int {
	numbers := []int{a, b, c, d, e} // initialize a slice to to store the input integers.
	n := len(numbers)               // Calculates the length of the numbers slice and stores it in the variable n.

	for i := 0; i < n-1; i++ { // bubble sort
		for j := i + 1; j < n-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return (numbers[n/2])
}
