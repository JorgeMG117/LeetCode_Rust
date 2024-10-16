package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func levelOrderRec(level int, root *TreeNode, values *[][]int) {
    // Add node to list of level
    if root == nil {
        return
    }
     
    if len(*values) == level {
        *values = append(*values, []int{root.Val})
    } else {
        (*values)[level] = append((*values)[level], root.Val)
    }

    fmt.Println(values)

    // Call to left with level++
    levelOrderRec(level+1, root.Left, values)
    // Call to right with level++
    levelOrderRec(level+1, root.Right, values)
}

func levelOrder(root *TreeNode) [][]int {
    vals := make([][]int, 0, 2000)

    levelOrderRec(0, root, &vals)
    
    return vals 
}


/*
Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/
func main() {
    // Tree 1: root = [3,9,20,null,null,15,7]
    root1 := &TreeNode{Val: 3}
    root1.Left = &TreeNode{Val: 9}
    root1.Right = &TreeNode{Val: 20}
    root1.Right.Left = &TreeNode{Val: 15}
    root1.Right.Right = &TreeNode{Val: 7}


    // Test the function with both trees
    fmt.Println("levelOrder: ", levelOrder(root1))
}

