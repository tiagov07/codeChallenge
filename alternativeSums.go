// Several people are standing in a row and need to be divided into two teams. The first person goes into
// team 1, the second goes into team 2, the third goes into team 1 again, the fourth into team 2, and so on.

// You are given an array of positive integers - the weights of the people. Return an array of two integers,
// where the first element is the total weight of team 1, and the second element is the total weight of
// team 2 after the division is complete.

func solution(a []int) []int {
	odd := 0
	pair := 0
	for i, val := range a {
		if i%2 == 0 {
			pair += val
		} else {
			odd += val
		}
	}

	result := make([]int, 0)
	result = append(result, pair)
	result = append(result, odd)

	return result
}

//alternative solution
func solution(a []int) []int {
	r := make([]int, 2)
	for i, w := range a {
		r[i%2] += w
	}
	return r
}