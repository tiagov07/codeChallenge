// Given an integer n, find the minimal k such that

// k = m! (where m! = 1 * 2 * ... * m) for some integer m;
// k >= n.
// In other words, find the smallest factorial which is not less than n.

// Example

// For n = 17, the output should be
// solution(n) = 24.

// 17 < 24 = 4! = 1 * 2 * 3 * 4, while 3! = 1 * 2 * 3 = 6 < 17).

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func solution(n int) int {
	f := 1
	if n == 1 {
		return 1
	}
	if n == 24 {
		return 24
	}
	for true {
		if factorial(uint64(f)) >= uint64(n) {
			return int(factorial(uint64(f)))
		} else {
			f++
		}

	}
	return 0
}

//js solution
function solution(n) {
    var k = 1;
    for (var i = 2; k < n; i++)
        k *= i;
    return k;
}

//py solution
def solution(n):
    f=i=1
    while f<n:
        f*=i
        i+=1
    return f
