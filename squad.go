package piscine

import "fmt"

// function QuadA to QuadE
// ////////////////////////
func QuadA(x, y int) {
	ln := y   //how many lines we will print
	char := x //how many characters we will print

	for i := 0; i < ln; i++ {
		if i == 0 {
			PrintFirstln(char, "o", "o", "-")
		} else if i == ln-1 {
			PrintLastln(char, "o", "o", "-")
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			PrintMidln(char, "|")
		}
		fmt.Print(("\n"))
	}

}

func QuadB(x, y int) {
	ln := y   //how many lines we will print
	char := x //how many characters we will print

	for i := 0; i < ln; i++ {
		if i == 0 {
			PrintFirstln(char, "/", "\\", "*")
		} else if i == ln-1 {
			PrintLastln(char, "\\", "/", "*")
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			PrintMidln(char, "*")
		}
		fmt.Print(("\n"))
	}

}

func QuadC(x, y int) {
	ln := y   //how many lines we will print
	char := x //how many characters we will print

	for i := 0; i < ln; i++ {
		if i == 0 {
			PrintFirstln(char, "A", "A", "B")
		} else if i == ln-1 {
			PrintLastln(char, "C", "C", "B")
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			PrintMidln(char, "B")
		}
		fmt.Print(("\n"))
	}

}

func QuadD(x, y int) {
	ln := y   //how many lines we will print
	char := x //how many characters we will print

	for i := 0; i < ln; i++ {
		if i == 0 {
			PrintFirstln(char, "A", "C", "B")
		} else if i == ln-1 {
			PrintLastln(char, "A", "C", "B")
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			PrintMidln(char, "B")
		}
		fmt.Print(("\n"))
	}

}

func QuadE(x, y int) {
	ln := y   //how many lines we will print
	char := x //how many characters we will print

	for i := 0; i < ln; i++ {
		if i == 0 {
			PrintFirstln(char, "A", "C", "B")
		} else if i == ln-1 {
			PrintLastln(char, "C", "A", "B")
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			PrintMidln(char, "B")
		}
		fmt.Print(("\n"))
	}

}

// we execute these fuctions when we call the QuadA, B, C, D, & E
// the for-loop function to print the first line
func PrintFirstln(char int, c1 string, c2 string, top string) {
	for i := 0; i < char; i++ {
		if i == 0 {
			fmt.Print(c1)
		} else if i == char-1 {
			fmt.Print(c2)
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			fmt.Print(top)
		}
	}
}

// the for-loop function to print the last line
func PrintLastln(char int, c3 string, c4 string, top string) {
	for i := 0; i < char; i++ {
		if i == 0 {
			fmt.Print(c3)
		} else if i == char-1 {
			fmt.Print(c4)
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			fmt.Print(top)
		}
	}
}

// the for-loop function to print the middle line(s)
func PrintMidln(char int, rue string) {
	for i := 0; i < char; i++ {
		if i == 0 {
			fmt.Print(rue)
		} else if i == char-1 {
			fmt.Print(rue)
		} else { // you can also add "if i != 0 || i != ln-1" but it's redundant
			fmt.Print(" ")
		}
	}
}
