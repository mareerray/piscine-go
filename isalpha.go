package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if (char >= 'a' && char <= 'z') && (char >= 'A' && char <= 'Z') && (char >= '0' && char <= '9') {
			return true
		} else if char == ' ' {
			return false
		}
	}
	return len(s) > 0
}

/* function that returns true if the string passed as the parameter
only contains alphanumerical characters
or is empty, otherwise, and returns false.*/
