package main

import "fmt"

type Pilot struct {
	Name     string
	Life     float64
	Age      int
	Aircraft string
}

func main() {

	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

const AIRCRAFT1 string = "1"
