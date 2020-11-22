package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := genListNode([]int{8})
	b := genListNode([]int{2})
	n := addTwoNumbers(a, b)
	printListNode(n)

	x := 3 % 10
	fmt.Println(x)
}

func genListNode(nums []int) *ListNode {
	res := &ListNode{}
	tmp := res
	for i, n := range nums {
		tmp.Val = n
		if i != len(nums)-1 {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}
	return res
}

func printListNode(n *ListNode) {
	res := []int{}
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}

	fmt.Println(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		val1, val2 := 0, 0

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (val1 + val2 + carry) % 10, Next: nil}
		current = current.Next
		carry = (val1 + val2 + carry) / 10

	}

	return head.Next
}
