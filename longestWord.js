// Define a word as a sequence of consecutive English letters. Find the longest word from the given string.

// Example

// For text = "Ready, steady, go!", the output should be
// solution(text) = "steady".
function solution(text) {
    var arr = text.match(/[a-z]+/gi);
       var sorted = arr.sort(function(a, b) {
       return b.length - a.length;
     });
     return sorted[0];
   }


//alternative solution
function solution(text) {
    
    let patt = /[^a-z^A-Z]/
    
    text = text.split(patt);
    
    let longest = text.reduce(function (a, b) { return a.length > b.length ? a : b; });
    
    return longest

}

//go solution
// import "regexp" 

// func solution(text string) string {
//     longest := ""
//     re := regexp.MustCompile("[A-Za-z]+")
//     for _, elem := range re.FindAllString(text, -1) {
//         if len(elem) > len(longest) {
//             longest = elem
//         }
//     }
//     return longest
// }