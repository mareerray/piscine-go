package piscine

func Divmod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}
