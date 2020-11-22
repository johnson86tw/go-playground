package main

import "math"

// TreeNode is ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return isValidbts(root, math.Inf(1), math.Inf(-1))
}

func isValidbts(root *TreeNode, max, min float64) bool {
	if root == nil {
		return true
	}

	v := float64(root.Val)
	return v < max && v > min && isValidbts(root.Left, v, min) && isValidbts(root.Right, max, v)
}
