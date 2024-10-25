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

// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/description/
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	tmp, pre := head, newHead
	for tmp != nil {
		for tmp.Next != nil && tmp.Val == tmp.Next.Val {
			tmp = tmp.Next
		}
		if pre.Next != tmp {
			pre.Next = tmp.Next
		} else {
			pre = tmp
		}
		tmp = pre.Next
	}
	return newHead.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
