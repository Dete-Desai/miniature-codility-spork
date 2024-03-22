package main

import "fmt"

func maxSubArraySum(array []int, size int) int {
	var max_ending_here int = 0
	var max_so_far int = 0

	for i := 0; i < size; i++ {
		max_ending_here = max_so_far + array[i]

		if max_ending_here < 0 {
			max_ending_here = 0
		}

		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
	}
	return max_so_far
}

func main() {
	fmt.Printf("The value of the maximum sub array sum is:")
}
