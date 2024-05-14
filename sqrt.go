package piscine

func Sqrt(nb int) int {
	if nb == 0 { // if function if nb equals to 0, return 0
		return 0
	}

	// goes through loop to find i where i*i equals to nb
	// stop looking for i if i*i is bigger than nb
	for i := 1; i*i <= nb; i++ {
		if i*i == nb { // if i*i is equal to nb
			return i // return answer
		}
	}
	return 0 // if function reaches here means to valid whole number
}

/*
Goal of the task: function to return the squar root of the int
passed as a parameter if that square root is a whole number.
Otherwise it returns 0.
*/
