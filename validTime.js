// Check if the given string is a correct time representation of the 24-hour clock.

// Example

// For time = "13:58", the output should be
// solution(time) = true;
// For time = "25:51", the output should be
// solution(time) = false;
// For time = "02:76", the output should be
// solution(time) = false.
function solution(time) {
    let parts = time.split(":")
    if(parts[0] < 24 && parts[1] < 60) return true
    else return false
}

//go solution
// import "strings"

// func solution(time string) bool {
//      parts := strings.Split(time, ":")
//      hours,_ := strconv.Atoi(parts[0])
//      minutes,_ := strconv.Atoi(parts[1])
//      if hours < 24 && minutes < 60 {return true} 
//      return false
// }

//alternative solution how to pass the string to a int variable
// func solution(time string) bool {
//     h,m:=0,0
//     fmt.Sscanf(time,"%d:%d",&h,&m)
//     return h < 24 && m < 60    
// }