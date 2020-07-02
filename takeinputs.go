package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//var myString string
	//fmt.Scanln(&myString)
	//fmt.Println(myString)

	//var name string = "hitesh"
	//var a, _ = fmt.Println(name)
	//fmt.Println(a)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your rating : ")
	//myname, _ := reader.ReadString('\n')
	//fmt.Println(myname)

	myrating, _ := reader.ReadString('\n')
	mynumrating, _ := strconv.ParseFloat(strings.TrimSpace(myrating), 64)
	fmt.Println(mynumrating + 2)

}
