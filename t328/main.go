package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// ol为奇数节点的尾部
	ol := head
	// e 为偶数节点的头部， el 为偶数节点的尾部
	e, el := head.Next, head.Next
	for ol.Next != nil && el.Next != nil {
		ol.Next = el.Next
		ol = ol.Next
		el.Next = ol.Next
		el = el.Next
	}
	ol.Next = e
	return head
}
