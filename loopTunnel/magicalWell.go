// You are standing at a magical well. It has two positive integers written on it:
// a and b. Each time you cast a magic marble into the well, it gives you a * b
// dollars and then both a and b increase by 1. You have n magic marbles.
// How much money will you make?

// Example

// For a = 1, b = 2, and n = 2, the output should be
// solution(a, b, n) = 8.

// You will cast your first marble and get $2, after which the numbers will
// become 2 and 3. When you cast your second marble, the well will give you $6.
// Overall, you'll make $8. So, the output is 8.

func solution(a int, b int, n int) int {
	cont := 0
	for i := 0; i < n; i++ {
		cont += a * b
		a += 1
		b += 1
	}
	return cont
}

//py solution
def solution(a, b, n):
    return sum([(a + i) * (b + i) for i in range(n)])

//js solution
function solution(a, b, n) {
    return n != 0 ? a * b + solution(++a, ++b, --n) : 0;
}