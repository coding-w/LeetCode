package main

import "fmt"

func main() {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}
	list2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	list3 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  6,
			Next: nil,
		},
	}
	lists := []*ListNode{list1, list2, list3}
	res := mergeKLists(lists)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}

// https://leetcode.cn/problems/merge-k-sorted-lists/description/
func mergeKLists(lists []*ListNode) *ListNode {
	var res *ListNode
	if len(lists) == 0 {
		return nil
	}
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		res = recursion(res, lists[i])
	}
	return res
}

func recursion(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = recursion(l1.Next, l2)
		return l1
	}
	l2.Next = recursion(l1, l2.Next)
	return l2
}

type ListNode struct {
	Val  int
	Next *ListNode
}
