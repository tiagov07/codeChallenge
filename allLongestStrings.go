// Given an array of strings, return another array containing all of its longest strings.

// Example

// For inputArray = ["aba", "aa", "ad", "vcd", "aba"], the output should be
// solution(inputArray) = ["aba", "vcd", "aba"].
package main

func solution(inputArray []string) []string {
	longestWord := 0
	solution := make([]string, 0)

	for _, el := range inputArray {
		if len(el) == longestWord {
			solution = append(solution, el)
		}
		if len(el) > longestWord {
			longestWord = len(el)
			solution = []string{el}

		}
	}
	return solution
}

//9
