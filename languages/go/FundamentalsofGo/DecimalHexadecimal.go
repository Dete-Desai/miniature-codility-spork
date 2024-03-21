package main

import "fmt"

func main() {
	adams := 50
	fmt.Printf("50 as a binary is %b\n", adams)
	fmt.Printf("50 as a hexadecimal is %x\n", adams)

	a, b, c, d, e, f := 0, 1, 2, 3, 4, 5

	fmt.Println(" Type | Default | Binary | Hexadecimal")
	fmt.Printf("   %T |      %v |     %b |         %#X \n", a, a, a, a)
	fmt.Printf("   %T |      %v |     %b |         %#X \n", b, b, b, b)
	fmt.Printf("   %T |      %v |     %b |         %#X \n", c, c, c, c)
	fmt.Printf("   %T |      %v |     %b |         %#X \n", d, d, d, d)
	fmt.Printf("   %T |      %v |     %b |         %#X \n", e, e, e, e)
	fmt.Printf("   %T |      %v |     %b |         %#X \n", f, f, f, f)
}
