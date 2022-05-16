// An IP address is a numerical label assigned to each device
// (e.g., computer, printer) participating in a computer network that uses the
// Internet Protocol for communication. There are two versions of the Internet
// protocol, and thus two versions of addresses. One of them is the IPv4 address.

// Given a string, find out if it satisfies the IPv4 address naming rules.

import "net"

func solution(inputString string) bool {
	return net.ParseIP(inputString) != nil
}

// import "strings"

// func solution(inputString string) bool {
//     adresses := strings.Split(inputString, ".")

//     if(len(adresses) != 4) {
//         return false
//     }

//     for i := range adresses {
//         adress, err := strconv.Atoi(adresses[i])
//         if(err != nil || adresses[i] == "" ||!(adress >= 0 && adress <= 255)) {
//             return false
//         }
//     }

//     return true
// }