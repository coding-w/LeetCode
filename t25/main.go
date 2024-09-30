package main

import "fmt"

func main() {
	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}},
			},
		},
	}
	res := reverseKGroup(list, 3)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next

	}

}

// https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p1, p2 := head, head
	t := 1
	for p2.Next != nil && t != k {
		p2 = p2.Next
		t++
	}
	if t < k {
		return p1
	} else {
		next := p2.Next
		p2.Next = nil
		reverse(p1)
		p1.Next = reverseKGroup(next, k)
		return p2
	}
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return node
}

type ListNode struct {
	Val  int
	Next *ListNode
}
