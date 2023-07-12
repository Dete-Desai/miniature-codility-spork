//When two or more conseccutive function parameters share a type,
//You can omit the type from all but the last.

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	a := add(100, 200)
	fmt.Println(a)
	fmt.Println(add(42, 13))
}
