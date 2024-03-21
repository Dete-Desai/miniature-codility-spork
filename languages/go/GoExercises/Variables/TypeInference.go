//This when you declare a variable without specifying the type using the := sign

package main

import "fmt"

func main() {
	i := 45
	y := float32(i) + float32(0.2)
	fmt.Printf("The value is:%v\n", y)
}
