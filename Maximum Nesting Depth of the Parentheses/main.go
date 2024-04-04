func maxDepth(s string) int {
    if len(s) < 2 {
        return 0
    }
    count := 0
    max := 0
    for _, char := range s {
        tmp := string(char)
        if tmp == "(" {
            count++
            if count > max {
                max = count
            }
        } else if tmp == ")" {
            count--
        }
    }
    return max
}