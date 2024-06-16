package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    firstOdd := &head
    firstEven := head.Next

    odd := firstOdd

    for odd != nil {

    }
    /*
    iter := head
    odd := true

    var changeOdd *ListNode
    var changeEven *ListNode

    for iter != nil {
        if odd {
            changeOdd = iter
        } else {
            nextEven
        }
    }

    odd := addEventList(iter.Ne
    */
    
    return head
}

func printLinkedList(head *ListNode) {
    iter := head
    for iter != nil {
        fmt.Print(iter.Val)
        fmt.Print(", ")
        iter = iter.Next
    }
    fmt.Println()
}

/*
Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
*/
func main() {
    input := []int{2,1,3,5,6,4,7}

    var list *ListNode
    for i := len(input)-1; i >= 0; i-- {
        v := input[i]
        node := &ListNode{
            Val: v,
            Next: list,
        }
        list = node
    }

    fmt.Println("Initial list")
    printLinkedList(list)
    result := oddEvenList(list)
    fmt.Println("Computed list")
    printLinkedList(result)
}
