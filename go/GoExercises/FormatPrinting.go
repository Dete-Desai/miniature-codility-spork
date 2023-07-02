package main

import "fmt"

func main() {
	const name, age = "Lamech", 28

	fmt.Printf("%s is %d years of age\nThe type of variables is %T for name and %T for age\n", name, age, name, age)
}
