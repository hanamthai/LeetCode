/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getValueByIndex(head *ListNode, idx int) int {
    current := head
    for i:=0;i<idx;i++ {
        current = current.Next
    }
    return current.Val
}

func reorderList(head *ListNode)  {
    if head == nil || head.Next == nil {
        return
    }
    length := 0
    var ruleValues []int
    current := head
    for current != nil {
        length++
        current = current.Next
    }
    start := 0
    end := length - 1 
    for start < end {
        ruleValues = append(ruleValues, start)
        ruleValues = append(ruleValues, end)
        start++
        end--
    }
    if length % 2 != 0 {
        ruleValues = append(ruleValues, start)
    }
    currentNewList := &ListNode{Val: head.Val}
    headCurrent := currentNewList
    for i, val := range ruleValues {
        if i == 0 {
            continue
        }
        value := getValueByIndex(head, val)
        newNode := &ListNode{Val: value}
        currentNewList.Next = newNode
        currentNewList = currentNewList.Next
    }
    head.Next = headCurrent.Next
}


// other way to resolver problem
func reorderList(head *ListNode)  {
    if head == nil{
        return
    }

    list := make([]*ListNode, 0)
    tmp := head

    for tmp != nil{
        list = append(list, tmp)
        tmp = tmp.Next
    }

    left, right := 0, len(list)-1
    for left < right{
        list[right].Next = nil
        list[left].Next = list[right]
        left++

        list[right].Next = list[left]
        list[left].Next = nil
        right--
    }
    head = list[0]
}