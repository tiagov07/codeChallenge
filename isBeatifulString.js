// A string is said to be beautiful if each letter in the string appears at most as many times as the previous letter in the alphabet within the
// string; ie: b occurs no more times than a; c occurs no more times than b; etc. Note that letter a has no previous letter.

// Given a string, check whether it is beautiful.

// Example

// For inputString = "bbbaacdafe", the output should be solution(inputString) = true

function solution(inputString) {
    return inputString.split('').sort().join('').replace(/([a-z])\1*/g,(it)=>it.length +',').split(',').slice(0,-1).map(Number).every((el,i,ar)=> (inputString.indexOf(String.fromCharCode(i+'a'.charCodeAt(0)))>=0 &&(i==0 || ar[i]<=ar[i-1])))
    }

// alternative solution
function solution(inputString) {
    s = "abcdefghijklmnopqrstuvwxyz"
    for (i = 1; i < s.length; i++) {
        if (inputString.split(s[i]).length-1 > inputString.split(s[i-1]).length-1) {
            return false}}
    return true}

//go solution
// func solution(inputString string) bool {
//     alpha := [26]rune{}
//       for _, v := range inputString {
//           c := v - 'a'
//           alpha[c]++
//       }
  
//       for i := 1; i < len(alpha); i++ {
//           if alpha[i] > alpha[i-1] {
//               return false
//           }
//       }
//       return true
//   }


// alternative
// func solution(inputString string) bool {
//     count := make([]int, 26)
//     for _, r := range inputString {
//         count[r - 'a']++
//     }
//     prev := 100
//     for _, c := range count {
//         if c > prev {
//             return false
//         }
//         prev = c
//     }
//     return true
// }
