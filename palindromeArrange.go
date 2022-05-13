// Given a string, find out if its characters can be rearranged to form a palindrome.

// Example

// For inputString = "aabb", the output should be
// solution(inputString) = true.

// We can rearrange "aabb" to make "abba", which is a palindrome.

func solution(inputString string) bool {
	var chars = make(map[rune]int)
	count := 0
	for _, val := range inputString {
		chars[val] += 1
		if chars[val]%2 == 0 {
			count -= 1
		} else {
			count += 1
		}
	}

	return count < 2
}