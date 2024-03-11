// https://leetcode.com/problems/maximum-odd-binary-number/?envType=daily-question&envId=2024-03-01

func maximumOddBinaryNumber(s string) string {
    // Hint: For a number to be odd, the last bit must be 1, and for a number to be the largest, 
    // all the remaining set bits must be on the left.
    lenString := len(s)
    freqOfOne := 0
    for _, v := range s {
        if string(v) == "1" {
            freqOfOne++
        }
    }
    var result string
    for i := 0; (freqOfOne - 1) > i;i++ {
        result+= "1"
    }
    for i := 0; (lenString - freqOfOne) > i;i++ {
        result+="0"
    }
    result+="1"
    return result
}