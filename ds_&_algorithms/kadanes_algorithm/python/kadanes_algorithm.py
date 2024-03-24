def maxSubarraySum(array):
    max_ending_here = max_so_far = 0

    for value in array:
        max_ending_here += value

        if(max_ending_here < 0):
            max_ending_here = 0
        max_so_far = max(max_so_far, max_ending_here)

    return max_so_far

if __name__ == "__main__":
    input_array = [2, 1, -3, 4, -1, 2, 1, -5, 4]
    max_sum = maxSubarraySum(input_array)
    print(f"The sum of the maximum sub arrays in the given array is: {max_sum}")