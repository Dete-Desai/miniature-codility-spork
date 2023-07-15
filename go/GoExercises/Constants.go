//Constants are declared like variables, but with the cost keyword
//Constants can be any basic types in golang
//Constants cannot be declared using the := (inference) syntax

package main

import "fmt"

const p = 3.14

func main() {
	const language = "##%#%%*(^&%^)"
	const validity = true

	fmt.Printf("Language is: %v\nPie is: %v\nTruth to the answer is: %v\n", language, p, validity)
}
