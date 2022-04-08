// Given the positions of a white bishop and a black pawn on the standard chess board, determine whether the bishop can
// capture the pawn in one move.

// The bishop has no restrictions in distance for each move, but is limited to diagonal movement. Check out the example
// below to see how it can move

// For bishop = "a1" and pawn = "c3", the output should be
// solution(bishop, pawn) = true.

function solution(bishop, pawn) {
    let board = {
        "a": 1,
        "b": 2,
        "c": 3,
        "d": 4,
        "e": 5,
        "f": 6,
        "g": 7,
        "h": 8
    }
    let bishopx = board[bishop[0]]
    let bishopy = parseInt(bishop[1])
    let pawnx = board[pawn[0]]
    let pawny = parseInt(pawn[1])
    
    if(bishopx + bishopy === pawny + pawnx || bishopx + pawny === bishopy + pawnx ){
        return true
    }
    return false
}

//alternative solution
function solution(bishop, pawn) {
    return Math.abs(bishop[0].charCodeAt()-pawn[0].charCodeAt())===Math.abs(bishop[1]-pawn[1])
  }

// in go
// func solution(bishop string, pawn string) bool {
//     x, y := bishop[0] - pawn[0], bishop[1] - pawn[1]
//     return x == y || x == -y
// }