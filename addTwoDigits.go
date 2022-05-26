// You are given a two-digit integer n. Return the sum of its digits.

// Example

// For n = 29, the output should be
// solution(n) = 11.

function solution(n) {
    return JSON.stringify(n).split("").map( el => parseInt(el) ).reduce( (ac, el ) => ac + el)
}

//alternative solution
function solution(n) {
    return n%10 + Math.floor(n/10)
}

//go solution
func solution(n int) int {
    firstDig := (n / 10)
    secondDig := n % 10
    return firstDig + secondDig
}

import "strings"

func solution(n int) int {
    letterParse := strconv.Itoa(n)
    letters := strings.Split(letterParse, "")
    result:= 0
    for _, val := range letters {
        number, _ := strconv.Atoi(val)
        result += number
    }
    return result
}



func solution(n int) int {
    digits := strconv.Itoa(n)
    left, _ := strconv.Atoi(string(digits[0]))
    right, _ := strconv.Atoi(string(digits[1]))
    return left + right
}