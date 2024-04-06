package main

import "fmt"

func sum(x int, y int) int {
	var result int = x + y // Constant time operation

	return result
}

func main() {
	var adder int = 10
	var addee int = 10

	var finalSum int = addee + adder

	fmt.Printf("The value of sum is %v", finalSum)
}
