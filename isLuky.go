// Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of
//the digits is equal to the sum of the second half.

// Given a ticket number n, determine if it's lucky or not.

// Example

// For n = 1230, the output should be
// solution(n) = true;
// For n = 239017, the output should be
// solution(n) = false.

func solution(n int) bool {
	parts := fmt.Sprintf("%v", n)
	middle := len(parts) / 2
	solution := 0

	for _, el := range parts[:middle] {
		solution += int(el - '0')
	}
	for _, el := range parts[middle:] {
		solution -= int(el - '0')
	}

	return solution == 0
}