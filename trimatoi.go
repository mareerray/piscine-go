package piscine

func TrimAtoi(s string) int {
	num := 0
	sign := 1
	started := false

	for _, ch := range s {
		if ch == '-' && !started {
			sign = -1
			started = true
		} else if ch >= '0' && ch <= '9' {
			num = num*10 + int(ch-'0')
			started = true
		}
	}

	return num * sign
}
