// Given an array of integers, find the maximal absolute difference between any two of its adjacent elements.

// Example

// For inputArray = [2, 4, 1, 0], the output should be
// solution(inputArray) = 3.

import "math"

func solution(inputArray []int) int {
	var maxAdjacent float64 = 0

	for i := 1; i < len(inputArray); i++ {
		if math.Abs(float64(inputArray[i-1]-inputArray[i])) > maxAdjacent {
			maxAdjacent = math.Abs(float64(inputArray[i-1] - inputArray[i]))
		}
	}
	return int(maxAdjacent)
}