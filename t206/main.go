package main

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	reverseList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/reverse-linked-list/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	res := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return res
}
