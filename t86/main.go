package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{2, nil}}}}}}
	res := partition(node, 3)
	printNode(res)
}

func printNode(res *ListNode) {
	tmp := res
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

// https://leetcode.cn/problems/partition-list/
func partition(head *ListNode, x int) *ListNode {
	small, bigger := &ListNode{0, nil}, &ListNode{0, nil}
	tmp1, tmp2 := small, bigger
	for head != nil {
		if head.Val >= x {
			tmp2.Next = &ListNode{head.Val, nil}
			tmp2 = tmp2.Next
		} else {

			tmp1.Next = &ListNode{head.Val, nil}
			tmp1 = tmp1.Next
		}
		head = head.Next
	}
	tmp1.Next = bigger.Next
	return small.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
