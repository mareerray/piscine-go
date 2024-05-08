// package is a container of source code for some specific purpose.
// the main package holds the code that is responsible to make the current program compilable&runnable
package main

import "github.com/01-edu/z01"

//Importing simply means bringing the specified package from its source location to the destination code.

/*func main() {
	z01.PrintRune('m')
	z01.PrintRune('\n')
}*/

/*func main() {
	for
	z01.PrintRune(rune(i))
}*/

func main() {
	for i := 97; i <= 122; i++ {
		z01.PrintRune(rune(i))

	}
	z01.PrintRune('\n')
}
