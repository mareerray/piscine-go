package piscine

func ConcatParams(args []string) string {
	var test string
	for i, arg := range args {
		test += arg
		if i < len(args)-1 {
			test += "\n"
		}
	}
	return test
}
