https://leetcode.com/problems/two-sum/

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			return []int{j, i} // Return the indices if complement exists in the map
		}
		numMap[num] = i // Store the current number along with its index in the map
	}

	return nil // If no solution is found, return nil
}

func main() {
	var nums = []int{2, 4, 6, 1, 9}
	var target int = 11
	fmt.Println("Result: ", twoSum(nums, target))
}
