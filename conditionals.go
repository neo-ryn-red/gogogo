package main

import (
	"fmt"
)

func main() {
	var isLoggedIn bool = true
	var balance int = 10

	if isLoggedIn || balance > 15 {
		fmt.Println("Show cart page")
		var len, err = fmt.Println("Show cart page\n")
		if err == nil {
			fmt.Printf("length is %v, %T\n", len, len)
			fmt.Println(`I am Backtick`)
		}

	} else {
		fmt.Println("Show user login page")
	}
}