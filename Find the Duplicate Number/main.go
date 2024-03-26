func findDuplicate(nums []int) int {
    if len(nums) == 0 {
        return -1
    }
    mapValWithBool := make(map[int]bool)
    for _, val := range nums {
        if mapValWithBool[val] == false {
            mapValWithBool[val] = true
        } else {
            return val
        }
    }
    return -1
}