// Given a rectangular matrix of characters, add a border of asterisks(*) to it.

// Example

// For

// picture = ["abc",
//            "ded"]
// the output should be

// solution(picture) = ["*****",
//                       "*abc*",
//                       "*ded*",
//                       "*****"]

import "strings"

func solution(picture []string) []string {
	border := make([]string, len(picture)+2)
	border[0] = strings.Repeat("*", len(picture[0])+2)
	border[len(border)-1] = strings.Repeat("*", len(picture[0])+2)
	for i := 1; i <= len(picture); i++ {
		border[i] = "*" + picture[i-1] + "*"
	}
	return border
}