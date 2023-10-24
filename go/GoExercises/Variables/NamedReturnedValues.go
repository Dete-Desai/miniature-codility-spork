//Go's returned values maybe named.
//If so they are treated as variables defined at the top of the function
//These names should be used to document the meaning of the return values.
//A return statement without arguments returns the named returned values also known as naked return
//Naked return statements sheould be used onl in short functions, as with the examplenshown here.
//They can harm readability in longer functions

package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
