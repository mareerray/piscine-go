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

// Starts the outer loop of the bubble sort algorithm.
// The loop iterates over each element in the numbers slice except for the last one.
// The -1 ensures that the last element is not compared with itself or with the last unsorted element.
// Begins the inner loop, which iterates through the unsorted part of the numbers slice.
// The expression n-i-1 calculates the number of remaining elements that need to be checked for sorting.
// As the outer loop progresses, fewer elements remain unsorted, hence the decrement.
// Checks if the current element (numbers[j]) is greater than the next element (numbers[j+1]).
// If true, it swaps these two elements. This is the core operation of the bubble sort algorithm,
// where larger elements "bubble up" towards the end of the slice.
// The expression n/2 calculates the index of the middle element,
// assuming n is the total number of elements in the slice.
