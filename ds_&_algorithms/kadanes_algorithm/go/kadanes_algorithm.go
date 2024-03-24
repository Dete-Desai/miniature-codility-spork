package main

import "fmt"

func maxSubArraySum(array []int) int {
	var max_ending_here int = 0
	var max_so_far int = 0

	for _, value := range array {
		max_ending_here += value

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
	var input_array []int = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	var max_sum int = maxSubArraySum(input_array)

	fmt.Printf("The maximum sum of sub arrays in the give array is %v", max_sum)
}
