// https://leetcode.com/problems/sqrtx/description/
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func mySqrt(x int) int {
	xSqrt := math.Sqrt(float64(x))
	result := int(xSqrt)
	return result
}

func main() {
	fmt.Println("Hello, 世界", mySqrt(8))
}
