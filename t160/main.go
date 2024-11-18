package main

func main() {

}

// https://leetcode.cn/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	//intersectVal := 0
	//listA := headA
	//listB := headB
	//skipA, skipB := 0, 0
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}
		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

type ListNode struct {
	Val  int
	Next *ListNode
}
