package piscine

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, char := range s {
		if (char >= 0 && char <= 31) || char == 127 {
			return false
		}
	}
	return true
}

/*Write a function that returns true if the string passed as a parameter
contains only printable characters, otherwise, returns false.*/
