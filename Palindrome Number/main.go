// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	var i, j int = 0, len(xStr) - 1
	for i < j {
		if xStr[i] != xStr[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	var x int = 121
	fmt.Println("Hello, 世界", isPalindrome(x))
}
