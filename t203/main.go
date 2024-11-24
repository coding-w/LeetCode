package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/remove-linked-list-elements/
func removeElements(head *ListNode, val int) *ListNode {
	res := &ListNode{Next: head}
	pre := res
	for head != nil {
		if head.Val == val {
			head = head.Next
			continue
		}
		pre.Next = head
		head = head.Next
		pre = pre.Next
	}
	pre.Next = head
	return res.Next
}
