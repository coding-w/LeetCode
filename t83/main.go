package main

import "fmt"

func main() {
	node := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}}}
	res := deleteDuplicates(node)
	printNode(res)
}

func printNode(res *ListNode) {
	tmp := res
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/description/
func deleteDuplicates(head *ListNode) *ListNode {
	tmp := head
	for tmp != nil {
		if tmp.Next != nil && tmp.Val == tmp.Next.Val {
			tmp.Next = tmp.Next.Next
			continue
		}
		tmp = tmp.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}
