package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	queue := []int{}
	queue = append(queue, 2)

	queue = queue[1:]
	fmt.Println(queue)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{}
	queue = append(queue, root)

	curNum, nextLevelNum := 1, 0
	res, tmp := [][]int{}, []int{}

	for len(queue) > 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}

			tmp = append(tmp, node.Val)
			queue = queue[1:]
			curNum--
		}

		if curNum == 0 {
			curNum = nextLevelNum
			nextLevelNum = 0
			res = append(res, tmp)
			tmp = []int{}
		}
	}

	return res
}
