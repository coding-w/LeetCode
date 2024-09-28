package main

import "fmt"

func main() {
	aa := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	res := removeNthFromEnd(aa, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	res := head
	recursion(res, n)
	return res
}

func recursion(head *ListNode, n int) int {
	if head == nil {
		return 0
	}
	index := recursion(head.Next, n)
	if n == index {
		head.Next = head.Next.Next
	}
	return index + 1
}

type ListNode struct {
	Val  int
	Next *ListNode
}
