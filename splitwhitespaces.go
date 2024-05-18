package piscine

func SplitWhiteSpaces(s string) []string {
	var singleword []string
	var word []rune

	for _, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if len(word) > 0 {
				singleword = append(singleword, string(word))
				word = []rune{}
			}
		} else {
			word = append(word, char)
		}
	}
	if len(word) > 0 {
		singleword = append(singleword, string(word))
	}
	return singleword
}
