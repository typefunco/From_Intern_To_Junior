package main

import "fmt"

func main() {
	variable1 := 5 // Just variable
	fmt.Println(variable1)

	variable2 := &variable1 // Variable address
	fmt.Println(variable2)

	variable3 := *variable2 // Just variable
	fmt.Println(variable3)
}
