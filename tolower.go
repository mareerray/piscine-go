package piscine

func ToLower(s string) string {
	var lowercase string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			lowercase = lowercase + string(char+32)
		} else {
			lowercase = lowercase + string(char)
		}
	}
	return lowercase
}
