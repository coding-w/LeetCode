package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	tmp := make([]int, 0)
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		tmp = append(tmp, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		tmp = append(tmp, slow.Val)
	}
	slow = slow.Next
	for i := len(tmp) - 1; i >= 0; i-- {
		if tmp[i] != slow.Val {
			return false
		}
		slow = slow.Next
	}
	return true
}
