package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	var firstOdd *ListNode
	var firstEven *ListNode
	isOdd := true

	var nextOdd **ListNode
	var nextEven **ListNode

	current := head

	for current != nil {
		if isOdd {
			//fmt.Println("Val is odd ", current.Val)
			isOdd = false

			if firstOdd == nil {
				firstOdd = current
				nextOdd = &current.Next
			} else {
				*nextOdd = current
				nextOdd = &current.Next
			}

		} else {
			//fmt.Println("Val is even ", current.Val)
			if firstEven == nil {
				firstEven = current
				nextEven = &current.Next
			} else {
				*nextEven = current
				nextEven = &current.Next
			}

			isOdd = true
		}

		next := current.Next
		current.Next = nil
		current = next

	}

	if nextOdd != nil {
		*nextOdd = firstEven
	}

	//fmt.Println((**nextOdd).Val)
	//fmt.Println((*firstEven).Val)
	//fmt.Println((*firstOdd).Val)

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
	input := []int{2, 1, 3, 5, 6, 4, 7}

	var list *ListNode
	for i := len(input) - 1; i >= 0; i-- {
		v := input[i]
		node := &ListNode{
			Val:  v,
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
