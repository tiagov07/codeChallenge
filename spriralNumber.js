// Construct a square matrix with a size N × N containing integers from 1 to N * N in a spiral order,
// starting from top-left and in clockwise direction.

// Example

// For n = 3, the output should be

// solution(n) = [[1, 2, 3],
//                     [8, 9, 4],
//                     [7, 6, 5]]

function solution(n) {
    var matrix = [...Array(n)].map(_ => []),
        x = 0,
        y = 0,
        dir = 2,
        size = n,
        c = 0;
    for (var i = 1; i <= n * n; i++) {
        matrix[y][x] = i;
        if (++c == size) {
            dir = (dir + 1) % 4;
            size -= dir % 2;
            c = 0;
        }
        if (dir % 2 == 0) x += dir > 1 ? 1 : -1;
        else y += dir > 1 ? 1 : -1;
    }
    return matrix;
}

//go solution
// func solution(n int) [][]int {
//     ret := make([][]int, n)
//     for i := range ret {
//         ret[i] = make([]int, n)
//     }
    
//     i := 0
//     j := 0
//     dirX := 1
//     dirY := 0
    
//     for count := 1; count <= n*n; count++ {
//         ret[j][i] = count
        
//         if !isValid(ret, i+dirX, j+dirY) {
//             if dirX == 0 {
//                 dirX = dirX - dirY
//                 dirY = 0
//             } else {
//                 dirY = dirX - dirY
//                 dirX = 0
//             }
//         }
        
//         i += dirX
//         j += dirY
//     }
    
//     return ret
// }

// func isValid(matrix [][]int, x, y int) bool {
//     if x >= len(matrix) || x < 0 || y >= len(matrix[0]) || y < 0 || matrix[y][x] != 0 {
//         return false
//     }
    
//     return true
// }