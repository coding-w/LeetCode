package main

import "fmt"

func main() {
	head := &ListNode{Val: 1,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 3,
				Next: &ListNode{Val: 4,
					Next: &ListNode{Val: 5, Next: nil}}}}}
	res := rotateRight(head, 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

// https://leetcode.cn/problems/rotate-list/description/
// tips: 链表转数组，数组旋转
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}
	// 链表 转 数组
	heads := make([]*ListNode, 0)
	for head != nil {
		heads = append(heads, &ListNode{Val: head.Val})
		head = head.Next
	}
	if len(heads) == 0 {
		return nil
	}
	length := len(heads)
	temp := make([]*ListNode, length)
	for i := 0; i < length; i++ {
		temp[(i+k)%length] = heads[i]
	}
	res := new(ListNode)
	item := new(ListNode)
	res.Next = item
	for i := 0; i < length; i++ {
		item.Val = temp[i].Val
		if i == length-1 {
			item.Next = nil
			break
		}
		item.Next = new(ListNode)
		item = item.Next
	}
	return res.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}
