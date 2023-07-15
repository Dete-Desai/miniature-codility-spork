//Numeric constants are high precision values
//An untyped constatnt takes the type needed by its constant

package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float32) float64 { return 1 * 0.1 }

func main() {
	fmt.Printf("The valule of Small when int is %v\n The value of Small when float is %v\n The value of Big when float is %v\n", needInt(Small), needFloat(Small), needFloat(Big))
}
