package main

func main() {
	res := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}
	reorderList(res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.cn/problems/reorder-list/
// tips: 快慢指针，慢指针找到中间点
func reorderList(head *ListNode) {
	if head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 后部 node
	tailNode := slow.Next
	slow.Next = nil
	// 反转 tailNode
	tailRes := (*ListNode)(nil)
	for tailNode != nil {
		p := tailNode.Next
		tailNode.Next = tailRes
		tailRes = tailNode
		tailNode = p
	}
	tmp := head
	for tmp != nil && tailRes != nil {
		r := tailRes
		tailRes = tailRes.Next
		l := tmp.Next
		tmp.Next = r
		r.Next = l
		tmp = l
	}
	return
}
