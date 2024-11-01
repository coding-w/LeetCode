package main

func main() {

}

// https://leetcode.cn/problems/same-tree/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSameTree(root.Left, root.Right)
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Right) && isSameTree(p.Right, q.Left)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
