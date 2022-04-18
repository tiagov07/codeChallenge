// Determine if the given character is a digit or not.

// Example

// For symbol = '0', the output should be
// solution(symbol) = true;
// For symbol = '-', the output should be
// solution(symbol) = false.

function solution(symbol) {
    numbers = [1,2,3,4,5,6,7,8,9,0]
    return numbers.some(x => x == symbol)
}

//alternative solution
function solution(symbol) {
    return !isNaN(symbol)}

//alternative
function solution(symbol) {
    return /\d/.test(symbol);
}

//go solution
// import "unicode"

// func solution(symbol string) bool {
//     return unicode.IsDigit(rune(symbol[0]))
// }