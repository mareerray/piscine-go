package piscine

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}

	for _, char := range s {
		if !(char >= 48 && char <= 57) {
			return false
		}
	}
	return true
}
