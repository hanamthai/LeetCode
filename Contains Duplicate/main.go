func containsDuplicate(nums []int) bool {
    mapValWithBool := make(map[int]bool)
    for _, num := range nums {
        if mapValWithBool[num] {
            return true
        }
        mapValWithBool[num] = true
    }
    return false
}