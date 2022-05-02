// Some people are standing in a row in a park. There are trees between them which cannot be moved.
// Your task is to rearrange the people by their heights in a non-descending order without moving the trees.
// People can be very tall!

// Example

// For a = [-1, 150, 190, 170, -1, -1, 160, 180], the output should be
// solution(a) = [-1, 150, 160, 170, -1, -1, 180, 190].

import "sort"

func solution(a []int) []int {
	arr2 := make([]int, 0)
	ind := make([]int, 0)
	for i, val := range a {
		if val != -1 {
			ind = append(ind, i)
			arr2 = append(arr2, val)
		}
	}
	sort.Ints(arr2)
	for i, val := range ind {
		a[val] = arr2[i]
	}

	return a
}