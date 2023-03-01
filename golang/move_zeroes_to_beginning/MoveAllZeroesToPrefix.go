package main

import "fmt"

/**
* given an input array, move all non-zero integers to suffix maintaining the original order or
* move all zeroes to the beginning of the array, leaving the non-zero elements in same mutual positions.
* Below is an O(n) time, O(1) space complexity solution.
 */
func main() {
	input := []int64{
		1, 0, 3, 4, 0, 6, 7, 0, 19,
	}
	printNumbersZeroPrefixed(input)
}

func printNumbersZeroPrefixed(input []int64) {
	lastZeroIndex := -1
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 0 {
			lastZeroIndex = i
			break
		}
	}

	for i := len(input) - 2; i >= 0; i-- {
		if input[i] > 0 && lastZeroIndex > 0 {
			input[lastZeroIndex] = input[i]
			lastZeroIndex--
			input[i] = 0
		}
	}

	for _, v := range input {
		fmt.Print(v, " ")
	}
}
