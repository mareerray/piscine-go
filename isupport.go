package piscine

func IsUpper(s string) bool {
	for _, char := range s { // iterates through each character in the string s.
		if char < 'A' || char > 'Z' { // Checks if the current character is not an uppercase letter.
			return false // If a non-uppercase character is found, the function returns false.
		}
	}
	return len(s) > 0 // After checking all characters, the function returns true only if the string is non-empty and all characters are uppercase.
}
