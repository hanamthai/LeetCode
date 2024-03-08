// You can edit this code!
// Click here and start typing.
// https://leetcode.com/problems/count-elements-with-maximum-frequency/description/?envType=daily-question&envId=2024-03-08
package main

import "fmt"

func maxFrequencyElements(nums []int) int {
	mapValueWithFrequency := make(map[int]int)
	for _, num := range nums {
		if _, ok := mapValueWithFrequency[num]; ok {
			mapValueWithFrequency[num]++
		} else {
			mapValueWithFrequency[num] = 1
		}
	}
	var totalFrequency, maxFreq int
	for _, frequency := range mapValueWithFrequency {
		if maxFreq < frequency {
			maxFreq = frequency
		}
	}
	for _, frequency := range mapValueWithFrequency {
		if maxFreq == frequency {
			totalFrequency += frequency
		}
	}
	return totalFrequency
}

func main() {
	a := []int{1, 2, 2, 3, 1, 4}
	fmt.Println("Hello, 世界", maxFrequencyElements(a))
}
