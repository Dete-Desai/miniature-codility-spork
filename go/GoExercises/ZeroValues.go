//Zero values are variables declared without an explicit value

package main

import "fmt"

func main() {
	var i int
	var b bool
	var s string
	var f float32

	fmt.Printf("%v\t%v\t%v\t%v\n", i, b, s, f)
}
