package piscine

var anyTrue bool //// Global variable to track if any string is numeric

func Any(f func(string) bool, a []string) bool {
	// passing item from the slice through the fucntion IsNumeric
	for _, value := range a {
		if f(value) {
			anyTrue = true
			break //// No need to check further elements since we found one that matches
		}
	}
	return anyTrue
}

/*
package piscine
func IsNumeric(s string) bool {
 if len(s) == 0 {
 return false
 }
 for _, char := range s {
 if !(char >= 48 && char <= 57) {
 return false
 }
 }
 return true
}
*/
