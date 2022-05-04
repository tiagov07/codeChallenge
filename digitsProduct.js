// Given an integer product, find the smallest positive (i.e. greater than 0) integer the product of
// whose digits is equal to product. If there is no such integer, return -1 instead.

// Example

// For product = 12, the output should be
// solution(product) = 26;
// For product = 19, the output should be
// solution(product) = -1.
function solution(product) {
    if(product == 0) return 10
    if(product === 1) return 1
    digits = ''
    for (let i = 9 ; i > 1; i --){
        while (product % i == 0) {
            product /= i
            digits = `${i}${digits}`
        }
    }
    if (product > 1 ) return -1
    return Number(digits)
}

//go solution
// func solution(product int) int {
//     if product == 0 {
//         return 10
//     }
//     if product == 1 {
//         return 1
//     }
    
//     digits := []int{}
//     fmt.Println(digits)
//     for i := 9; i > 1; i-- {
//         for (product % i) == 0{
//             digits = append(digits,i)
//             product /= i
//         }
//     }
//     if product != 1 {
//         return -1
//     }
//     total := 0
//     for i := len(digits)-1; i >= 0; i-- {
//        total = 10*total + digits[i] 
//     }
//     fmt.Println(digits)
//     return total
// }