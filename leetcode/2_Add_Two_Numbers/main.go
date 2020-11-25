package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	n := addTwoNumbers(genListNode([]int{2, 4, 3}), genListNode([]int{5, 6, 4}))
	printListNode(n)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	tmp := res
	adding := 0

	for l1 != nil || l2 != nil {
		l1val, l2val := 0, 0

		if l1 != nil {
			l1val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2val = l2.Val
			l2 = l2.Next
		}

		sum := l1val + l2val + adding

		if sum/10 > 0 {
			adding = sum / 10
			sum = sum % 10
		} else {
			adding = 0
		}

		tmp.Val = sum

		if l1 == nil && l2 == nil && adding == 0 {
			break
		}

		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}

	return res
}

func genListNode(nums []int) *ListNode {
	res := &ListNode{}
	tmp := res
	for i, n := range nums {
		tmp.Val = n
		if i == len(nums)-1 {
			break
		}
		tmp.Next = &ListNode{}
		tmp = tmp.Next
	}
	return res
}

func printListNode(l *ListNode) {
	res := []int{}
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	fmt.Println(res)
}
