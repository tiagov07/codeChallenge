// Two arrays are called similar if one can be obtained from another by
// swapping at most one pair of elements in one of the arrays.

// Given two arrays a and b, check whether they are similar.

// Example

// For a = [1, 2, 3] and b = [1, 2, 3], the output should be
// solution(a, b) = true.

// The arrays are equal, no need to swap any elements.

// For a = [1, 2, 3] and b = [2, 1, 3], the output should be
// solution(a, b) = true.

// We can obtain b from a by swapping 2 and 1 in b.

// For a = [1, 2, 2] and b = [2, 1, 1], the output should be
// solution(a, b) = false.

// Any swap of any two elements either in a or in b won't
// make a and b equal.

func solution(a []int, b []int) bool {
	aDiff := make([]int, 0)
	bDiff := make([]int, 0)
	cont := 0

	for i, _ := range a {
		if a[i] != b[i] {
			aDiff = append(aDiff, a[i])
			bDiff = append(bDiff, b[i])
			cont++
		}
	}
	if cont == 0 {
		return true
	}

	if cont == 1 || cont > 2 {
		return false
	}

	if aDiff[0] == bDiff[1] && aDiff[1] == bDiff[0] {
		return true
	}

	return false

}

//alternative solution
func solution(a []int, b []int) bool {
	c := 0
	s := make(map[int]bool)
	for n, _ := range a {
		if a[n] != b[n] {
			c++
			s[a[n]] = true
			s[b[n]] = true
		}

	}

	return (c == 2 || c == 0) && (len(s) == 0 || len(s) == 2)
}

//16