package main

import "fmt"

type ByteSize = int

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	EB
	ZB
)

func main() {
	fmt.Printf("%d \t %b \n", KB, KB)
	fmt.Printf("%d \t %b \n", MB, MB)
	fmt.Printf("%d \t %b \n", GB, GB)
	fmt.Printf("%d \t %b \n", TB, TB)
	fmt.Printf("%d \t %b \n", EB, EB)
	fmt.Printf("%d \t %b \n", ZB, ZB)
}
