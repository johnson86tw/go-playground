package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Parent struct {
	Val   int
	Exist bool
}

func main() {

}

func isValidBST(root *TreeNode) bool {
	p := Parent{Val: 0, Exist: false}
	return isValid(root, p)
}

func isValid(root *TreeNode, p Parent) bool {
	if root == nil {
		return true
	}

	if root.Right != nil {
		if root.Val < root.Right.Val {
			if p.Exist && root.Right.Val < p.Val {
				return true
			}
		}
	}

	if root.Left != nil {
		if root.Val > root.Left.Val {
			if p.Exist && root.Left.Val > p.Val {
				return true
			}
		}
	}

	p = Parent{Val: root.Val, Exist: true}

	return isValid(root.Left, p) && isValid(root.Right, p)

}
