package piscine

func Abort(a, b, c, d, e int) int {
	sortedSet := []int{a, b, c, d, e}
	for i := 0; i < len(sortedSet)-1; i++ {
		for j := i + 1; j < len(sortedSet); j++ {
			if sortedSet[i] > sortedSet[j] {
				sortedSet[i], sortedSet[j] = sortedSet[j], sortedSet[i]
			}
		}
	}
	return sortedSet[2]
}
