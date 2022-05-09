// You are given an array of integers. On each move you are allowed to increase exactly one of its element by one. Find the minimal number of
// moves required to obtain a strictly increasing sequence from the input.

// Example

// For inputArray = [1, 1, 1], the output should be
// solution(inputArray) = 3.

func solution(inputArray []int) int {
	moves := 0

	for i := 1; i < len(inputArray); i++ {
		if d := inputArray[i-1] - inputArray[i]; d >= 0 {
			inputArray[i] += d + 1
			moves += d + 1
		}
	}
	return moves
}