// https://leetcode.com/problems/longest-common-prefix/description/

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	minLen := len(strs[0])
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	prefix := make([]byte, minLen)
	for i := 0; i < minLen; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[0][i] {
				return string(prefix[:i])
			}
		}
		prefix[i] = strs[0][i]
	}

	return string(prefix)
}

func main() {
	a := []string{"flower", "flow", "flight"}
	fmt.Println("Hello, 世界", longestCommonPrefix(a))
}
