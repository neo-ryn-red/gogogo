package main

import "fmt"

func main() {
	var batman string = "I am batman"
	fmt.Println(batman)

	var superman string = "I am superman"
	fmt.Println(superman)

	thor := "I am thor"
	fmt.Println(thor)

	//thorrating := 3.
	//fmt.Printf("%v, %T", thorrating, thorrating)

	var Ironman, capAmerica string = "I am Ironman", "I am cap"
	fmt.Println(Ironman, capAmerica)

	var defaultValue string
	fmt.Println(defaultValue)

	var (
		spiderman = "I am spider"
		age       = 18
		powers    = "Web slings, spidy sense"
		antman    = "I am antman"
	)

	fmt.Println(spiderman, age, powers, antman)
}