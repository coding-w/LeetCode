package main

func main() {
	recoverTree(nil)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/recover-binary-search-tree/
func recoverTree(root *TreeNode) {
	var n1, n2, pre *TreeNode
	var recovered func(*TreeNode)
	recovered = func(head *TreeNode) {
		if head == nil {
			return
		}
		recovered(head.Left)
		if pre != nil && head.Val > pre.Val {
			n2 = head
			if n1 == nil {
				n1 = pre
			} else {
				return
			}
		}
		pre = head
		recovered(head.Right)

	}
	recovered(root)
	n1.Val, n2.Val = n2.Val, n1.Val
}
