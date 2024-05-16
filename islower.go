package piscine

func IsLower(s string) bool {
	for _, char := range s { // iterates through each character in the string s.
		if char < 'a' || char > 'z' { // Checks if the current character is not an lowercase letter.
			return false // If a non-lowercase character is found, the function returns false.
		}
	}
	return len(s) > 0 // After checking all characters, the function returns true only if the string is non-empty and all characters are uppercase.
}
