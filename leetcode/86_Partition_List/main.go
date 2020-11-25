package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// DoublyListNode ...
type DoublyListNode struct {
	Val  int
	Prev *DoublyListNode
	Next *DoublyListNode
}

func main() {
	// nums := []int{1, 4, 3, 2, 5, 2}
	// target := 3
}

// func foo(nums []int, target int) []int {
// 	for i, num := range nums {
// 		if num < target {
// 			prevs := nums[0:i]
// 			compareNum := num
// 			for j := len(prevs) - 1; j > -1; j-- {
// 				if prevs[j] >= target && prevs[j-1] < target {
// 					break
// 				}
// 				compareNum = prevs[j]
// 			}

// 		}
// 	}
// }

func genDoublyListNode(head *ListNode) *DoublyListNode {
	cur := head.Next
	DLNHead := &DoublyListNode{Val: head.Val, Prev: nil, Next: nil}
	curDLN := DLNHead
	for cur != nil {
		tmp := &DoublyListNode{Val: cur.Val, Prev: curDLN, Next: nil}
		curDLN.Next = tmp
		curDLN = tmp
		cur = cur.Next
	}
	return DLNHead
}

func genListNode(head *DoublyListNode) *ListNode {
	cur := head.Next
	LNHead := &ListNode{Val: head.Val, Next: nil}
	curLN := LNHead
	for cur != nil {
		tmp := &ListNode{Val: cur.Val, Next: nil}
		curLN.Next = tmp
		curLN = tmp
		cur = cur.Next
	}
	return LNHead
}

// 單向鏈表
func partition(head *ListNode, x int) *ListNode {
	left, right := &ListNode{}, &ListNode{}
	leftHead := left
	rightHead := right
	// 遞迴 head 比較 Val 與 x，分別丟進 left 和 right 裡頭
	for node := head; node != nil; node = node.Next {

		if node.Val < x {
			left.Next = node
			left = left.Next
		} else {
			right.Next = node
			right = right.Next
		}
	}

	right.Next = nil
	left.Next = rightHead.Next

	return leftHead.Next
}

func solution2(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	DLNHead := genDoublyListNode(head)
	cur := DLNHead
	for cur != nil {
		// 只有小於 target 的才進入
		if cur.Val < x {
			tmp := &DoublyListNode{Val: cur.Val, Prev: nil, Next: nil}
			// 預設為當前輪到的 node
			compareNode := cur
			// 重新定義 compareNode 為大於等於 target 的第一個 node
			for compareNode.Prev != nil {
				if compareNode.Val >= x && compareNode.Prev.Val < x {
					break
				}
				compareNode = compareNode.Prev
			}

			// 若 compareNode 是第一個 node
			if compareNode == DLNHead {
				if compareNode.Val < x {
					cur = cur.Next
					continue
					// 若第一個 node 大於 target
				} else {
					tmp.Next = DLNHead
					DLNHead.Prev = tmp
					DLNHead = tmp
				}
				// 若 compareNode 不是第一個 node
			} else {
				// 將 tmp 插在 compare 之前
				tmp.Next = compareNode
				tmp.Prev = compareNode.Prev
				compareNode.Prev.Next = tmp
				compareNode.Prev = tmp
			}
			// 開始刪除 node 的動作
			// 預設 deleteNode 為當前輪到的 node
			deleteNode := cur
			// cur 不是第一個 node
			if cur.Prev != nil {
				deleteNode.Prev.Next = deleteNode.Next
			}
			// cur 不是最後一個 node
			if cur.Next != nil {
				deleteNode.Next.Prev = deleteNode.Prev
			}
		}
		cur = cur.Next
	}
	return genListNode(DLNHead)
}
