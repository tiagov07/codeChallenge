// n children have got m pieces of candy. They want to eat as much candy as they can, but each child must
// eat exactly the same amount of candy as any other child. Determine how many pieces of candy will be eaten by
// all the children together. Individual pieces of candy cannot be split.

// Example

// For n = 3 and m = 10, the output should be
// solution(n, m) = 9.

// Each child will eat 3 pieces. So the answer is 9.

import "math"

func solution(n int, m int) int {
	division := math.Round(float64(m / n))
	fmt.Println(division)
	return int(division) * n
}

//alternative solution
func solution(n int, m int) int {
	return (m / n) * n
}

//alternative solution
func solution(n int, m int) int {
	return m - (m % n)
}

//js solution
function solution(n, m) {
    return m - m % n;
}

