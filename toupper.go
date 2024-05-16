package piscine

func ToUpper(s string) string {
	var uppercase string
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			uppercase = uppercase + string(char-32)
		} else {
			uppercase = uppercase + string(char)
		}
	}
	return uppercase
}
