// In order to stop the Mad Coder evil genius you need to decipher the
// encrypted message he sent to his minions.
// The message contains several numbers that,
// when typed into a supercomputer, will launch a missile into the sky
// blocking out the sun, and making all the people on Earth grumpy and
// sad.

// You figured out that some numbers have a modified single digit in their
// binary representation. More specifically, in the given number n the
// kth bit from the right was initially set to 0, but its current value
// might be different. It's now up to you to write a function that will
// change the kth bit of n back to 0.

// Example

// For n = 37 and k = 3, the output should be
// solution(n, k) = 33.

// 3710 = 1001012 ~> 1000012 = 3310.

// For n = 37 and k = 4, the output should be
// solution(n, k) = 37.

// The 4th bit is 0 already (looks like the Mad Coder forgot to encrypt
// 	this number), so the answer is still 37.

func solution(n int, k int) int {
	return n &^ (1 << uint(k-1))
}

//explanation
// The expression 1 << (k - 1) shifts the number 1 exactly k-1 times to
// the left so as an example for an 8 Bit number and k = 4:

// Before shift: 00000001
// After Shift: 00010000

// This marks the bit to kill. You see, 1 was shifted to the fourth
// position, as it was on position zero. The operator ^ negates each bit,
// meaning 1 becomes 0 and 0 becomes 1. For our example:

// Before negation: 00010000
// After negation: 11101111

// At last, & executes a bit-wise AND on two operands. Let us say,
// we have a number n = 17, which is 00010001 in binary. Our example
// now is:

// 00010001 & 11101111 = 00000001

// This is, because each bit of both numbers is compared by AND on
// the same position. Only positions, where both numbers have a 1
// remain 1, all others are set to 0. Consequently, only position
// zero remains 1.

// Overall your method int killKthBit(int n, int k) does exactly
// that with binary operators, it sets the bit on position k of number
// n to 0.