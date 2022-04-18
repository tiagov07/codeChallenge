// Given a string, return its encoding defined as follows:

// First, the string is divided into the least possible number of disjoint substrings
// consisting of identical characters
// for example, "aabbbc" is divided into ["aa", "bbb", "c"]
// Next, each substring with length greater than one is replaced with a concatenation of
// its length and the repeating character
// for example, substring "bbb" is replaced by "3b"
// Finally, all the new strings are concatenated together in the same order and a new
// string is returned.
function solution(s) {
    return s.replace(/(.)\1*/g, (chars, i) => (i === chars ? i : chars.length + i));
    }

//alternative solution
function solution(s) {
    let count = 1;
    let ans = '';
    
    for (let i = 0; i < s.length; i++) {
      if (s[i] === s[i + 1]) {
        count++;
      } else {
        if (count > 1) ans += count + s[i];
        else ans += s[i];
        count = 1;
      }
    }
    return ans;
  }

//go solution
// func solution(s string) string {
//     ans := ""
//     for i, count := 0, 0; i < len(s); i++ {
//         count++
//         if i+1 == len(s) || s[i] != s[i+1] {
//             if count > 1 {
//                 ans += fmt.Sprintf("%d", count)
//             }
//             ans += fmt.Sprintf("%c", s[i])
//             count = 0
//         }
//     }
//     return ans
// }