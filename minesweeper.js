// In the popular Minesweeper game you have a board with some mines and those cells that don't contain a mine have a number in it
// that indicates the total number of mines in the neighboring cells. Starting off with some arrangement of mines we want to create a Minesweeper
// game setup.

// Example

// For

// matrix = [[true, false, false],
//           [false, true, false],
//           [false, false, false]]
// the output should be

// solution(matrix) = [[1, 2, 1],
//                        [2, 1, 1],
//                        [1, 1, 1]]
function solution(matrix) {
    let height = matrix.length;
    let width = matrix[0].length;
    
    let outArray = Array.from(Array(height), () => new Array(width));
    
    let mines = 0;
    
    for(let i = 0; i < height; i++){
        for(let j = 0; j < width; j++){
            mines = 0;
            
            if(i>0){
                if(matrix[i-1][j-1]) mines += 1;
                if(matrix[i-1][j]) mines += 1;
                if(matrix[i-1][j+1]) mines += 1;  
            }
            
            if(i < height - 1){
                if(matrix[i+1][j-1]) mines += 1;
                if(matrix[i+1][j]) mines += 1;
                if(matrix[i+1][j+1]) mines += 1;
            }
            
            if(matrix[i][j-1]) mines += 1;
            if(matrix[i][j+1]) mines += 1;
            
            outArray[i][j] = mines;
        }
    }
    return outArray
}

func solution(matrix [][]bool) [][]int {
    m := make([][]int, len(matrix))
    for i, arr := range matrix {
        m[i] = make([]int, len(arr))
        for j, _ := range arr {
            for a:=i-1;a<=i+1;a++{
                if a < 0 || a >= len(matrix) {
                    continue
                }
                for b:=j-1;b<=j+1;b++{
                    if b < 0 || (a == i && b == j) || b >= len(arr) {
                        continue
                    }
                    if matrix[a][b] {
                        m[i][j]++
                    }
                }
            }
        }
    }
    return m
}

//24