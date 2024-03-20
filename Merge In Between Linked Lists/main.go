func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    cloneNode := list1
    current := cloneNode
    for i:=0; i<a-1;i++ {
        current = current.Next
    }
    endNode := current
    for i:=0; i<=(b-a)+1;i++ {
        endNode = endNode.Next
    }
    current.Next = list2
    for current.Next != nil {
        current = current.Next
    }
    current.Next = endNode
    return cloneNode
}