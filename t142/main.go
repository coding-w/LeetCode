package main

func main() {

}

// https://leetcode.cn/problems/linked-list-cycle-ii/description/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		if fast.Next.Next == slow.Next {
			break
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}

type ListNode struct {
	Val  int
	Next *ListNode
}
