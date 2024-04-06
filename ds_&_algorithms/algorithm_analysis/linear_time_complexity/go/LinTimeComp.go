package main

import "fmt"

func sumOperation(n int) int {
	var sum int = 0

	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	var value int = 3
	var finalSum = sumOperation(value)

	fmt.Printf("The value of the operation is: %v", finalSum)
}
