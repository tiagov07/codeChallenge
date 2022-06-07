// Given integers a and b, determine whether the following pseudocode results in an
// infinite loop

// while a is not equal to b do
//   increase a by 1
//   decrease b by 1
// Assume that the program is executed on a virtual machine which can store arbitrary
// long numbers and execute forever.

// Example

// For a = 2 and b = 6, the output should be
// solution(a, b) = false;
// For a = 2 and b = 3, the output should be
// solution(a, b) = true.

func solution(a int, b int) bool {
	if a%2 != b%2 || a > b {
		return true
	} else {
		return false
	}
}

// alternative solution
func solution(a int, b int) bool {
	return a > b || (b-a)%2 == 1
}
