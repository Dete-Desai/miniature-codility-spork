package main

import "fmt"

func recursive_function(n int) int {
	if n != 0 {
		return n * recursive_function(n-1)
	} else {
		return 1
	}
}

func main() {
	var number int = 4
	var result int = recursive_function(number)
	fmt.Printf("%d factorial = %d", number, result)
}
