// Given some integer, find the maximal number you can obtain by deleting exactly one digit of the given number.

// Example

// For n = 152, the output should be
// solution(n) = 52;
// For n = 1001, the output should be
// solution(n) = 101.
function solution(n) {
    s = n.toString()
    return Math.max(...[...Array(s.length).keys()].map(i=>Number(s.slice(0,i)+s.slice(i+1))))}

// go solution
// func solution(n int) int {
//     s := strconv.Itoa(n)
//     max := 0
//     for i := 0; i < len(s); i++ {
//         ds := string(s[0:i]) + string(s[i+1:])
//         dn, _ := strconv.Atoi(ds)
//         if dn > max {
//             max = dn
//         }
//     }
    
//     return max
// }