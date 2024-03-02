// You can edit this code!
// Click here and start typing.
// https://leetcode.com/problems/remove-element/
package main

import "fmt"

func removeElement(nums []int, val int) int {
	var list []int
	for _, num := range nums {
		if num != val {
			list = append(list, num)
		}
	}
	for i := range list {
		nums[i] = list[i]
	}
	return len(list)
}

func main() {
	var sliceTest = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println("Before change slice", sliceTest)
	fmt.Println("Hello, 世界", removeElement(sliceTest, 2))
	fmt.Println("After change slice", sliceTest)
}
