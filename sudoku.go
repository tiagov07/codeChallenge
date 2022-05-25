// Sudoku is a number-placement puzzle. The objective is to fill a 9 × 9 grid with digits so that each column, each row, and each of the nine 3 × 3
// sub-grids that compose the grid contains all of the digits from 1 to 9.

// This algorithm should check if the given grid of numbers represents a correct solution to Sudoku.

// Example

// For
// grid = [[1, 3, 2, 5, 4, 6, 9, 8, 7],
//         [4, 6, 5, 8, 7, 9, 3, 2, 1],
//         [7, 9, 8, 2, 1, 3, 6, 5, 4],
//         [9, 2, 1, 4, 3, 5, 8, 7, 6],
//         [3, 5, 4, 7, 6, 8, 2, 1, 9],
//         [6, 8, 7, 1, 9, 2, 5, 4, 3],
//         [5, 7, 6, 9, 8, 1, 4, 3, 2],
//         [2, 4, 3, 6, 5, 7, 1, 9, 8],
//         [8, 1, 9, 3, 2, 4, 7, 6, 5]]
// the output should be
// solution(grid) = true;

// For
// grid = [[1, 3, 2, 5, 4, 6, 9, 2, 7],
//         [4, 6, 5, 8, 7, 9, 3, 8, 1],
//         [7, 9, 8, 2, 1, 3, 6, 5, 4],
//         [9, 2, 1, 4, 3, 5, 8, 7, 6],
//         [3, 5, 4, 7, 6, 8, 2, 1, 9],
//         [6, 8, 7, 1, 9, 2, 5, 4, 3],
//         [5, 7, 6, 9, 8, 1, 4, 3, 2],
//         [2, 4, 3, 6, 5, 7, 1, 9, 8],
//         [8, 1, 9, 3, 2, 4, 7, 6, 5]]
// the output should be
// solution(grid) = false.

// The output should be false: each of the nine 3 × 3 sub-grids should contain all of the digits from 1 to 9.

function solution(grid) {
    for (let i = 0; i < 9; i ++) {
        let col = new Set();
        let row = new Set();
        for (let j = 0; j < 9; j ++) {
            row.add(grid[i][j]);
            col.add(grid[j][i]);
        }
        if (col.size < 9 || row.size < 9) {
            return false;
        }
    }
    
    for (let i = 0; i < 9; i += 3) {
        for (let j = 0; j < 9; j += 3) {
            let threeByThree = new Set();
            for (let k = 0; k <3 ; k ++) {
                for (let m = 0; m < 3; m ++) {
                    threeByThree.add(grid[i+k][j+m]);
                }
            }
            if (threeByThree.size < 9) {
                return false;
            }
        }
    }

    return true;
}

// go solution
// func solution(grid [][]int) bool {
//     if grid[0][0] == grid[0][1] {
//         return false
//     }
//     for y:=0; y<9;y++{
//         sy:=0
//         sx:=0
//         s:=0
//         for x:=0;x<9;x++{
//             sy+=grid[y][x]
//             sx+=grid[x][y]
//             s+=grid[y-y%3+x/3][y%3*3+x%3]
//         }
//         if sy!= 45 || sx != 45 || s!= 45 {
//             return false
//         }
//     }
//     return true
// }

