// https://leetcode.com/problems/length-of-last-word/
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func lengthOfLastWord(s string) int {
	var result int
	start := len(s) - 1
	for string(s[start]) == " " {
		start--
	}
	for i := start; i >= 0; i-- {
		if string(s[i]) == " " {
			break
		} else {
			result++
		}
	}
	return result
}

func main() {
	a := "   fly me   to   the moon  "
	fmt.Println("Hello, 世界", lengthOfLastWord(a))
}
