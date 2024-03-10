https://leetcode.com/problems/intersection-of-two-arrays/?envType=daily-question&envId=2024-03-10

func intersection(nums1 []int, nums2 []int) []int {
    var result []int
    if len(nums1) == 0 || len(nums2) == 0 {
        return result
    }
    var mapValueWithBool = make(map[int]bool)
    for _, num1 := range nums1 {
        mapValueWithBool[num1] = true
    }

    for _, num2 := range nums2 {
        if mapValueWithBool[num2] {
            result = append(result, num2)
            mapValueWithBool[num2] = false
        }
    }
    return result
}