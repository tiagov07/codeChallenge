// An email address such as "John.Smith@example.com" is made up of a local part ("John.Smith"), an "@" symbol, then a domain part ("example.com").

// The domain name part of an email address may only consist of letters, digits, hyphens and dots. The local part, however, also allows a lot of 
// different special characters. Here you can look at several examples of correct and incorrect email addresses.

// Given a valid email address, find its domain part.

function solution(address) {
    return /@([a-zA-Z0-9.\-]+)$/.exec(address)[1];
    }

//alternative solution
function solution(address) {
    return address.split('@').pop()
}

// go solution

// func solution(address string) string {
//     for i:=len(address)-1;i>0;i--{
//         if address[i] == '@' {
//             return address[i+1:]
//         }
//     }
//     return ""
// }

// import "strings"
// func solution(address string) string {
//     s := strings.Split(address, "@")
//     return s[len(s) - 1]
// }