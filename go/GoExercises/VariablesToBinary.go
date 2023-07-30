package main

import fmt "fmt"

/*
Printing verbs to decimal, binary and hexadecimal
*/

func main() {

	x, y, z := 747, 911, 90210

	fmt.Printf("The value of x\n as decimal is: %d,\n as a binary is: %b,\n as a hexadecimal is: %#X\n\n", x, x, x)
	fmt.Printf("The value of y\n as decimal is: %d,\n as a binary is: %b,\n as a hexadecimal is: %#X\n\n", y, y, y)
	fmt.Printf("The value of z\n as decimal is: %d,\n as a binary is: %b,\n as a hexadecimal is: %#X\n\n", z, z, z)
}
