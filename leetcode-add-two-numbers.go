package main

import (
    "fmt"
)

type ListNode struct {
    val  int
    Next *ListNode
}

func main() {
	
    l1 := &ListNode{3, &ListNode{4, &ListNode{5, nil}}}
    l2 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}

    result := addTwoNumbers(l1, l2)
    for node := result; node != nil; node = node.Next {
        fmt.Print(node.val)
    }
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyHead := &ListNode{0, nil}
    current := dummyHead
    carry := 0

    for l1 != nil || l2 != nil || carry > 0 {
        sum := carry
        if l1 != nil {
            sum += l1.val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.val
            l2 = l2.Next
        }

        carry = sum / 10
        current.Next = &ListNode{sum % 10, nil}
        current = current.Next
    }

    return dummyHead.Next
}
