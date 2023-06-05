package main

import (
	"fmt"
)

func main() {

	var batman string = " i am batman"
	fmt.Println(batman)

	var sup string
	sup = "i am superman"
	fmt.Println(sup)

	thor := "i am thor"
	fmt.Println(thor)

	thorRating := 3.
	fmt.Printf("%v, %T", thorRating, thorRating)

	fmt.Println()

	var score int = 4
	fmt.Printf("Hello world.. %v", score)

	fmt.Println()

	var Ironman, CatAmerica string = "I am Ironman", "I am capt America"
	fmt.Println(Ironman, CatAmerica)

	var defaultValue float64
	fmt.Println(defaultValue)

	var defaultValue1 int
	fmt.Println(defaultValue1)

	var defaultValue2 string
	fmt.Println(defaultValue2)

	var (
		sprintman = "I am spiderman "
		age       = 18
		powers    = "web slings, spidy sense"
		antman    = "i am antman"
	)

	fmt.Println(sprintman, age, powers, antman)
}
