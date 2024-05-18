package piscine

func AppendRange(min, max int) []int {
	var mintomax []int // mintomax is a variable type slice of int
	mintomaxLength := max - min

	if min >= max {
		return nil
	}

	for i := 0; i < mintomaxLength; i++ {
		mintomax = append(mintomax, min+i)
	}
	return mintomax
}
