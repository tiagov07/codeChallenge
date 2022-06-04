// Some phone usage rate may be described as follows:

// first minute of a call costs min1 cents,
// each minute from the 2nd up to 10th (inclusive) costs min2_10 cents
// each minute after 10th costs min11 cents.
// You have s cents on your account before the call. What is the duration of the longest call
// (in minutes rounded down to the nearest integer) you can have?

// Example

// For min1 = 3, min2_10 = 1, min11 = 2, and s = 20, the output should be
// solution(min1, min2_10, min11, s) = 14.

// Here's why:

// the first minute costs 3 cents, which leaves you with 20 - 3 = 17 cents;
// the total cost of minutes 2 through 10 is 1 * 9 = 9, so you can talk 9 more minutes and still have
// 17 - 9 = 8 cents;
// each next minute costs 2 cents, which means that you can talk 8 / 2 = 4 more minutes.
// Thus, the longest call you can make is 1 + 9 + 4 = 14 minutes long.
func solution(min1 int, min2_10 int, min11 int, s int) int {
	totalMinutes := 0
	if s > 0 && s >= min1 {
		s = s - min1
		totalMinutes += 1
	} else {
		return totalMinutes
	}
	for i := 0; i < 9; i++ {
		if s > 0 && s >= min2_10 {
			s = s - min2_10
			totalMinutes += 1
		} else {
			return totalMinutes
		}
	}
	totalMinutes += s / min11
	return totalMinutes
}

//alternative solution
func solution(min1 int, min2_10 int, min11 int, S int) int {
	if min1 == S {
		return 1
	}
	if min1 > S {
		return 0
	}
	if min1+(9*min2_10) > S {
		return (S-min1)/min2_10 + 1
	}
	return ((S - (min1 + (9 * min2_10))) / min11) + 10
}