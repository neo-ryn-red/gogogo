package main

import (
	"fmt"
)

func main() {
	var score int = 777
	p := &score

	if p != nil {
		fmt.Println("P is Having a value : ", *p)
	} else {
		fmt.Println("watchout for nil values")
	}
}
