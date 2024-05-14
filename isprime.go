package piscine

func IsPrime(nb int) bool {
	if nb <= 1 { // first check if nb is greater than or eqaul to 1
		return false // as 0 and 1 are not prime number
	}
	if nb == 2 {
		return true // 2 is the only even prime number
	}
	if nb%2 == 0 {
		return false // exclude even numbers greater than 2
	}
	for i := 3; i < nb; i++ {
		if nb%i == 0 {
			return false // nb is not prime
		}
	}
	return true
}

/*
Goal of the task: a function to return true if the int passed
as parameter is a prime number.
Otherwise it returns false.

A prime number is a number greater than 1 that can only be divided evenly
by two numbers: 1 and itself. This means it can't be made
by multiplying two smaller numbers together.
The function must be optimized in order to avoid time-outs with the tester.
(We consider that only positive numbers can be prime numbers)
(We also consider that 1 is not a prime number)

*/
