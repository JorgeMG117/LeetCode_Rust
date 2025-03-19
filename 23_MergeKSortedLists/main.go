package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
[
	1->4->5,
	1->3->4,
	2->6
]

1 1 2

3 4 6

4 5

No tiene sentido trabajar con cada lista (ya esta ordenada)
Tenemos que hacer algo que use varias listas
Opciones
La midad de las listas disponibles(no lo veo)
No tengo longitud de las cadenas, por lo que tampoco puedo partir estas por la mitad

Por lo tanto trabajamos con los primeros de cada cadena




merging them into one sorted list:
1->1->2->3->4->4->5->6
*/

func mergeKLists(lists []*ListNode) *ListNode {

	for i, v := range lists {
		fmt.Println(i, v.Val)
		//Create list
		//Order list
	}

	return lists[0]
}

func main() {
	values := [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}

	fmt.Println(values)

	var lists []*ListNode

	createList := func(nums []int) *ListNode {
		var head, tail *ListNode
		for _, num := range nums {
			newNode := &ListNode{Val: num}
			if head == nil {
				head = newNode
				tail = newNode
			} else {
				tail.Next = newNode
				tail = tail.Next
			}
		}
		return head
	}

	// Convert each slice of integers into a linked list
	for _, vals := range values {
		list := createList(vals)
		lists = append(lists, list)
	}

	// Merge all lists
	mergedList := mergeKLists(lists)

	// Print the merged list
	for node := mergedList; node != nil; node = node.Next {
		fmt.Print(node.Val, " ")
	}
	fmt.Println()

}
