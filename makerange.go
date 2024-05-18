package piscine

func MakeRange(min, max int) []int {
	if min >= max { // check if min is greater than or equal to max
		return nil
	}

	var mintomax []int
	mintomaxLength := max - min
	mintomax = make([]int, mintomaxLength)

	for i := 0; i < mintomaxLength; i++ {
		mintomax[i] = min + i
	}
	return mintomax
}
