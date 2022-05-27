// Given an integer n, return the largest number that contains exactly n digits.

// Example

// For n = 2, the output should be
// solution(n) = 99.

import (
	"strings"
)

func solution(n int) int {
	list := make([]string, 0)
	for i := 0; i < n; i++ {
		list = append(list, "9")
	}

	fmt.Println(list)
	result, _ := strconv.Atoi(strings.Join(list, ""))
	return result

}

//alternative solution
import "math"
func solution(n int) int {
    return int(math.Pow10(n)-1)

}

//alternative solution
func solution(n int) int {
    max := 1
    for i := 0; i < n; i++ {
        max *= 10
    }
    return max - 1
}



//js solution
function solution(n) {
    return Math.pow(10, n) - 1
}

//solution 2
function solution(n) {
    return Number('9'.repeat(n));
}