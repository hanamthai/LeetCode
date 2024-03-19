import "fmt"

func findMinArrowShots(points [][]int) int {
    if len(points) == 0 {
        return 0
    }
    sort.Slice(points, func(x, y int) bool {
        return points[x][1] < points[y][1]
    })
    result := 1
    current := points[0][1]
    for i:=1; i<len(points);i++ {
        if points[i][0] > current {
            current = points[i][1]
            result++
        }
    }
    return result
}
