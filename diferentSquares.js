// Given a rectangular matrix containing only digits, calculate the number of different 2 × 2 squares in it.

// Example

// For

// matrix = [[1, 2, 1],
//           [2, 2, 2],
//           [2, 2, 2],
//           [1, 2, 3],
//           [2, 2, 1]]
// the output should be
// solution(matrix) = 6.

// Here are all 6 different 2 × 2 squares

function solution(matrix) {
    s = new Set()
    for (i = 0; i < matrix.length-1; i++) {
        for (j = 0; j < matrix[i].length-1; j++) {
            s.add([matrix[i][j], matrix[i][j+1], matrix[i+1][j], matrix[i+1][j+1]].toString())}}
    return s.size
}

//go solution
// func solution(matrix [][]int) int {
//     m := make(map[string]int)
//     for y:=0; y<len(matrix) - 1; y++ {
//         for x:=0; x < len(matrix[y]) - 1; x++ {
//             k:=fmt.Sprintf("%v.%v.%v.%v", matrix[y][x], matrix[y+1][x], matrix[y][x+1], matrix[y+1][x+1])
//             m[k] ++
//         }
//     }
//     return len(m)
// }

//55