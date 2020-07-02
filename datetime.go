package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var name string
	var userrating string

	//front end
	//take name input
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your ful name : ")
	name, _ = reader.ReadString('\n')

	//take rating from user and conv to int / float
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our dosa center bet 1 and 5 : ")
	userrating, _ = reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(userrating), 64)

	//back end
	fmt.Printf("Hello %v\n, Thanks for rating our dosa center with %v star rating. \n\n Your rating was recorded in our system at %v\n\n", name, mynumrating, time.Now().Format(time.Stamp))

	if mynumrating == 5 {
		fmt.Println("Bonus for team for 5 star service")
	} else if mynumrating == 4 || mynumrating == 3 {
		fmt.Println("We are always improving")
	} else if mynumrating < 3 {
		fmt.Println("Need serious work on our side")
	}
}
