package main

import "fmt"

func main() {
	var firstName, secondName, age = "Lamech", "Dete", 28
	var height float32 = 6.3
	occupation, religion, weight, shoeSize := "Civil Engineer", "Christianity", 93, 8.5

	fmt.Printf("%s %s is %d years with a height of %v.", firstName, secondName, age, height)
	fmt.Printf("His religion is %s with a %s degree from Havard.", religion, occupation)
	fmt.Printf("%s likes jordan shoes with a shoe size of %v because of his weight of %d\n", firstName, shoeSize, weight)
}
