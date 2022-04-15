// Elections are in progress!

// Given an array of the numbers of votes given to each of the candidates so far, and an integer k equal to the number of voters who haven't
// cast their vote yet, find the number of candidates who still have a chance to win the election.

// The winner of the election must secure strictly more votes than any other candidate. If two or more candidates receive the same (maximum)
// number of votes, assume there is no winner at all.

// Example

// For votes = [2, 3, 5, 2] and k = 3, the output should be
// solution(votes, k) = 2.

function solution(votes, k) {
    let count = 0
    let firstOne = Math.max(...votes)
    let filter = votes.filter(x => x == firstOne)
    votes.map(x => x + k > firstOne ? count++ : count)
    if (k == 0 && filter.length == 1){
        count++
    }
    return count
}

//alternative solution
function solution(votes, k) {
    var max=Math.max(...votes)
    var r=votes.filter(x=>x+k>max||x===max).length
    return k?r:r==1?1:0
  }

//go solution
// func solution(votes []int, k int) int {
//     maxVote := 0
//     maxCnt := 0
    
//     for _, vote := range votes {
//         if vote == maxVote {
//             maxCnt++
//         } else if vote > maxVote {
//             maxVote = vote
//             maxCnt = 1
//         } 
//     }

//     result := 0
    
//     for _, vote := range votes {
//         //  fmt.Println(">", vote + k)
//         if vote + k > maxVote {
//             result ++
//         }
//     }
    
//     if maxCnt == 1 && k == 0 {
//         result ++
//     }
    
//     return result
    
// }