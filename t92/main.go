package main

func main() {

}

// https://leetcode.cn/problems/reverse-linked-list-ii/description/
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	start := pre.Next
	then := start.Next

	for i := 0; i < right-left; i++ {
		start.Next = then.Next
		then.Next = pre.Next
		pre.Next = then
		then = start.Next
	}

	return dummy.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
