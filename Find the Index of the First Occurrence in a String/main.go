// https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}
	lenNeedle := len(needle)
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+lenNeedle] == needle {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "hello"
	needle := "ll"
	fmt.Println("Hello, 世界", strStr(haystack, needle))
}
