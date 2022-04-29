// Given two strings, find the number of common characters between them.

// Example

// For s1 = "aabcc" and s2 = "adcaa", the output should be
// solution(s1, s2) = 3.

// Strings have 3 common characters - 2 "a"s and 1 "c".

func solution(s1 string, s2 string) int {
	cont := 0
	seen := make(map[rune]int)
	for _, el := range s1 {
		seen[el]++
	}
	for _, el := range s2 {
		if seen[el] > 0 {
			seen[el]--
			cont++
		}
	}
	return cont
}

//10