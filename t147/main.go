package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	res := new(ListNode)
	res.Next = head
	pre := res
	tmp := head
	for tmp != nil {
		next := tmp.Next
		if next != nil && next.Val < tmp.Val {
			for pre.Next.Val < next.Val {
				pre = pre.Next
			}
			tmp.Next = next.Next
			next.Next = pre.Next
			pre.Next = next
			pre = res
		} else {
			tmp = next
		}
	}
	return res.Next
}
