// https://leetcode.com/problems/search-insert-position/
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func searchInsertByBinarySearch(nums []int, target int) int {
	// binary search
	left, right := 0, len(nums)
	for left < right {
		mid := int((right + left) / 2)
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func searchInsertNormal(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	a := []int{1, 3, 5, 6}
	fmt.Println("Hello, 世界", searchInsertByBinarySearch(a, 5))
}
