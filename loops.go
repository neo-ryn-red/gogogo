package main

import (
	"fmt"

)

func main() {
	start := 1
	things := []string{"maleta", "bolso", "boleto", "vaso", "casa"}

	fmt.Println(things)

	for i := 0; i < 10; i++ {
		fmt.Println(i + start)
	}

	for i := 0; i < len(things); i++ {
		fmt.Println(things[i])
	}

	for i := range things {
		fmt.Println(things[i])
	}


	for start < 100 {
		start += start
		if start == 32 {
			goto lcolabel
		}
		fmt.Println("Start is now : ", start)
	}


	lcolabel: fmt.Println("My ass is nice")

}