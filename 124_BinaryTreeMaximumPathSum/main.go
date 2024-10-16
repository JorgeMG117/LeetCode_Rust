package main

import (
    "fmt"
)

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

const MIN = -1000

func maxOfSix(a, b, c, d, e, f int) int {
    max := a
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }
    if d > max {
        max = d
    }
    if e > max {
        max = e
    }
    if f > max {
        max = f
    }
    return max
}

func maxOfThree(a, b, c int) int {
    max := a
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }
    return max
}

// Return (current node path, best path so far)
func maxPathSumRec(root *TreeNode) (int, int) {
    l, r, b_l, b_r := 0, 0, MIN, MIN
    if root.Left != nil {
        l, b_l = maxPathSumRec(root.Left)
    }
    if root.Right != nil {
        r, b_r = maxPathSumRec(root.Right)
    }
    
    //v := root.Val + l + r

    // Select better val
    best := maxOfSix(root.Val + l + r, root.Val + r, root.Val + l, root.Val, b_l, b_r)

    // v should be max between
    // v + l or v + r or v
    v := maxOfThree(root.Val + r, root.Val + l, root.Val)

    // best should be max between 
    // maxOfSix


    //fmt.Println("Choosing: ", best)

    return v, best
}

func maxPathSum(root *TreeNode) int {
    _, best := maxPathSumRec(root)
    return best
}

/*
A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. Note that the path does not need to pass through the root.

The path sum of a path is the sum of the node's values in the path.

Given the root of a binary tree, return the maximum path sum of any non-empty path.
*/
func main() {
    // Tree 1: root = [1,2,3]
    root1 := &TreeNode{Val: 1}
    root1.Left = &TreeNode{Val: 2}
    root1.Right = &TreeNode{Val: 3}

    // Tree 2: root = [-10,9,20,null,null,15,7]
    root2 := &TreeNode{Val: -10}
    root2.Left = &TreeNode{Val: 9}
    root2.Right = &TreeNode{Val: 20}
    root2.Right.Left = &TreeNode{Val: 15}
    root2.Right.Right = &TreeNode{Val: 7}

    // Test the function with both trees
    fmt.Println("Max Path Sum of Tree 1:", maxPathSum(root1)) // Expected: 6 (2 -> 1 -> 3)
    fmt.Println("Max Path Sum of Tree 2:", maxPathSum(root2)) // Expected: 42 (-10 -> 20 -> 15 -> 7)

    // Tree: [5,4,8,11,null,13,4,7,2,null,null,null,1]
    root := &TreeNode{Val: 5}
    root.Left = &TreeNode{Val: 4}
    root.Right = &TreeNode{Val: 8}

    root.Left.Left = &TreeNode{Val: 11}
    root.Left.Left.Left = &TreeNode{Val: 7}
    root.Left.Left.Right = &TreeNode{Val: 2}

    root.Right.Left = &TreeNode{Val: 13}
    root.Right.Right = &TreeNode{Val: 4}
    root.Right.Right.Right = &TreeNode{Val: 1}

    // Test the maxPathSum function
    fmt.Println("Max Path Sum:", maxPathSum(root)) // Expected output: 48
}
