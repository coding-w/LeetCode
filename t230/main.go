package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	res := make([]int, 0)
	var recursion func(root *TreeNode)
	recursion = func(root *TreeNode) {
		if root == nil || len(res) == k {
			return
		}
		recursion(root.Left)
		res = append(res, root.Val)
		recursion(root.Right)
	}
	recursion(root)
	return res[k-1]
}
