package piscine

func Capitalize(s string) string {
	result := ""
	isNotCapital := true
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			if isNotCapital {
				result = result + string(char-'a'+'A')
				isNotCapital = false // is now capital
			} else {
				result = result + string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			if isNotCapital {
				result = result + string(char)
				isNotCapital = false // is now capital
			} else {
				result = result + string(char-'A'+'a')
			}
		} else if char >= '0' && char <= '9' {
			result = result + string(char)
			isNotCapital = false
		} else {
			result = result + string(char)
			isNotCapital = true
		}
	}
	return result
}
