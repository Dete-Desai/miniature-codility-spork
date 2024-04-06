package main

import "fmt"

func polTimeComp(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("i=%d, j=%d", i, j)
		}
		fmt.Println("End of inner loop")
	}
	fmt.Println("End of outer loop")
}

func main() {
	var value int = 3

	fmt.Println("The value of the loop is:\n")

	polTimeComp(value)
}
