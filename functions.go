package main

import (
	"fmt"

)

func main() {
	superman()
}

func superman(){
	fmt.Println("I am superman")
	result := multiplyme(3, 6)
	fmt.Println(result)
	myresult, mylength, myname := adder(3, 6, 7, 4, 7, 8)
	fmt.Println(myresult, mylength, myname)
}

func multiplyme(v1, v2 int) int {
	return v1 * v2
}

func adder(values ... int) (int, int, string) {
	sum := 0
	for i := range values{
		sum = sum +values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}