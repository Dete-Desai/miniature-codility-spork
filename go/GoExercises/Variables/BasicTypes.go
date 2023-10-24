package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint       = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 121)
)

func main() {
	fmt.Printf("%T\t%v\n", ToBe, ToBe)
	fmt.Printf("%T\t%v\n", MaxInt, MaxInt)
	fmt.Printf("%T\t%v\n", z, z)
}
