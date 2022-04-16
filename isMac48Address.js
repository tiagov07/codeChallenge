// A media access control address (MAC address) is a unique
// identifier assigned to network interfaces for communications on
// the physical network segment.

// The standard (IEEE 802) format for printing MAC-48 addresses in
// human-friendly form is six groups of two hexadecimal digits
// (0 to 9 or A to F), separated by hyphens
// (e.g. 01-23-45-67-89-AB).

// Your task is to check by given string inputString whether
// it corresponds to MAC-48 address or not.

function solution(inputString) {
    let parts = inputString.split("-")
    return parts.length == 6 && parts.every((a)=>/^[A-F0-9]{2}$/.test(a)) 
}

//aternative solution
solution = s => /^([0-9A-F]{2}-){5}[0-9A-F]{2}$/.test(s)


//go solution
// import "net"
// func solution(inputString string) bool {
//     _, err := net.ParseMAC(inputString)
//     return err == nil
// }

//go solution 2
// import "regexp"
// func solution(inputString string) bool {
//     re := regexp.MustCompile(`^([A-F0-9][A-F0-9]-){5}[A-F0-9][A-F0-9]$`)
//     return re.MatchString(inputString)
// }