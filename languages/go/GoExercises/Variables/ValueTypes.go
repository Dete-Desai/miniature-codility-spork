package main

import "fmt"

//Storing variables of different types

func main() {
	s, i, f := "Happiness", 42, 42.42
	fmt.Printf("The value of s is: %v and its type is: %T\n", s, s)
	fmt.Printf("The value of s is: %v and its type is: %T\n", i, i)
	fmt.Printf("The value of s is: %v and its type is: %T", f, f)
}
