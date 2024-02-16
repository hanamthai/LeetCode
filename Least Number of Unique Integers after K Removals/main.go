// https://leetcode.com/problems/least-number-of-unique-integers-after-k-removals/?envType=daily-question&envId=2024-02-16

func findLeastNumOfUniqueInts(arr []int, k int) int {
    mapIntWithFreq := make(map[int]int)
    for _, val := range arr {
        mapIntWithFreq[val]++
    }
    var freqArr []int
    for _, v := range mapIntWithFreq {
        freqArr = append(freqArr, v)
    }
    sort.Ints(freqArr)
    result := len(mapIntWithFreq)
    for _, freq := range freqArr {
        if k >= freq {
            k -= freq
            result -= 1
        } else {
            break
        }
    }
    return result
}