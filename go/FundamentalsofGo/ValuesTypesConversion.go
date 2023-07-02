package main

import "fmt"

func main() {
	age := 42
	weight := 93.0

	fmt.Printf("Default value of age is %v and its Type is %T\n", age, age)
	fmt.Printf("Default value of weight is %v and its Type is %T\n", weight, weight)

	var temperature float32 = 50.405

	fmt.Printf("Default value of temperature is %v and its Type is %T\n", temperature, temperature)

	weight = float64(temperature)

	fmt.Printf("Default value of weight is %v and its Type is %T\n", weight, weight)
}
