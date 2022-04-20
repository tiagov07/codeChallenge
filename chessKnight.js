// Given a position of a knight on the standard chessboard, find the number of different moves the knight can perform.

// The knight can move to a square that is two squares horizontally and one square vertically, 
// or two squares vertically and one square horizontally away from it. The complete move therefore looks
// like the letter L. Check out the image below to see all valid moves for a knight piece that is placed on
// one of the central squares.

function solution(cell) {
    const chessBoard = [
        [2,3,4,4,4,4,3,2],
        [3,4,6,6,6,6,4,3],
        [4,6,8,8,8,8,6,4],
        [4,6,8,8,8,8,6,4],
        [4,6,8,8,8,8,6,4],
        [4,6,8,8,8,8,6,4],
        [3,4,6,6,6,6,4,3],
        [2,3,4,4,4,4,3,2],
    ]
    
    const x = cell[0].charCodeAt() - 97;
    console.log(x)
    return chessBoard[cell[1] - 1][x]
}

// go solution
// func solution(cell string) int {
//     let := cell[0] 
//     num := cell[1] - '0'
    
//     ret := 8
//     if let == 'a' || let == 'h' {
//         ret = ret / 2
//     }
//     if let == 'b' || let == 'g' {
//         ret = ret - 2
//     }
//     if num == 1 || num == 8 {
//         ret = ret / 2
//     }
//     if num == 2 || num == 7 {
//         ret = ret - 2
//     }
//     return ret
// }