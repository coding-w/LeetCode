package main

func main() {

}

// https://leetcode.cn/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil {
		_, ok := m[head]
		if !ok {
			m[head] = true
		} else {
			return true
		}
		head = head.Next

	}
	return false

}

type ListNode struct {
	Val  int
	Next *ListNode
}
