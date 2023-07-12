//A function can return any number of results
//Using the swipe function we can return two strings

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "world")
	fmt.Println(a, b)
}
