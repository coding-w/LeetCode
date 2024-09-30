package main

import "fmt"

func main() {
	list := &ListNode{
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
	res := swapPairs(list)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

// https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := head.Next
	head.Next = swapPairs(p.Next)
	p.Next = head
	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}
