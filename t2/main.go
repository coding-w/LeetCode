package main

func main() {
	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 9}
	l1.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next = &ListNode{Val: 9}
	l1.Next.Next.Next.Next = &ListNode{Val: 9}
	l2 := &ListNode{Val: 9}
	l2.Next = &ListNode{Val: 9}
	l2.Next.Next = &ListNode{Val: 9}
	head := addTwoNumbers(l1, l2)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{Val: 0}
	temp := head
	for l1 != nil || l2 != nil {
		sum := temp.Val
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		temp.Val = sum % 10
		if !(sum/10 == 0 && (l1 == nil && l2 == nil)) {
			temp.Next = &ListNode{Val: sum / 10}
			temp = temp.Next
		}
	}
	return head
}
