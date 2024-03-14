func majorityElement(nums []int) int {
    conditionCmp := float64(len(nums)) / float64(2)
    mapValueWithFreq := make(map[int]int)
    for _, v := range nums {
        mapValueWithFreq[v] += 1
    }
    for v, freq := range mapValueWithFreq {
        if float64(freq) > conditionCmp {
            return v
        } 
    }
    return 0
}