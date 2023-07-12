//A var declaration can include initializers, one per variable.
//If an initializer is present, the type can be omitted; the variable will take the type of the intializer.

package main

import "fmt"

var i, j = 1, 2

func main() {
	var c, python, java = true, false, "no!"
	fmt.Printf("%T \t %T \t %T \t %T\t %T \n", i, j, c, python, java)
}
