// Call two people equally strong if their strongest arms are equally strong (the strongest arm can be both the right and the left),
// and so are their weakest arms.

// Given your and your friend's arms' lifting capabilities find out if you two are equally strong.

// Example

// For yourLeft = 10, yourRight = 15, friendsLeft = 15, and friendsRight = 10, the output should be
// solution(yourLeft, yourRight, friendsLeft, friendsRight) = true;
// For yourLeft = 15, yourRight = 10, friendsLeft = 15, and friendsRight = 10, the output should be
// solution(yourLeft, yourRight, friendsLeft, friendsRight) = true;
// For yourLeft = 15, yourRight = 10, friendsLeft = 15, and friendsRight = 9, the output should be
// solution(yourLeft, yourRight, friendsLeft, friendsRight) = false.

func solution(yourLeft int, yourRight int, friendsLeft int, friendsRight int) bool {
	yourMax := max(yourLeft, yourRight)
	yourMin := min(yourLeft, yourRight)

	fMax := max(friendsLeft, friendsRight)
	fMin := min(friendsLeft, friendsRight)

	return yourMax == fMax && yourMin == fMin
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//19