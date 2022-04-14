// Given a string, find the shortest possible string which can be achieved by adding characters to the end of initial string to make it a palindrome.

// Example

// For st = "abcdc", the output should be
// solution(st) = "abcdcba".

function solution(st) {
    var i = 0;
    var aux;
    while(st != st.split('').reverse().join('')){
        aux = st.split('')
        aux.splice(st.length-i, 0 ,st[i])
        st = aux.join('');
        i++;
    }
    return st;
}

//go solution
// func isPalindrome(st string) bool {
//     if len(st) == 0 || len(st) == 1 {
//         return true
//     }
//     if st[0] != st[len(st)-1] {
//         return false
//     }
//     return isPalindrome(st[1 : len(st)-1])
// }

// func solution(st string) string {
//     tailst := ""
//     for _, elem := range st {
//         if isPalindrome(st + tailst) {
//             return st + tailst
//         }
//         tailst =  string(elem) + tailst
//     }
//     if isPalindrome(st + tailst) {
//         return st + tailst
//     }
//     return ""
// }