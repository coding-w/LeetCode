package main

import "fmt"

func main() {
	l := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: &ListNode{Val: 3}}}}
	res := sortList(l)
	fmt.Println(res.Val)
}

// https://leetcode.cn/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	r := sortList(slow.Next)
	slow.Next = nil
	l := sortList(head)
	return merge(r, l)
}

func merge(r *ListNode, l *ListNode) *ListNode {
	res := &ListNode{}
	tmp := res
	for l != nil && r != nil {
		if l.Val <= r.Val {
			tmp.Next = l
			l = l.Next
		} else {
			tmp.Next = r
			r = r.Next
		}
		tmp = tmp.Next
	}
	if l != nil {
		tmp.Next = l
	} else {
		tmp.Next = r
	}
	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
