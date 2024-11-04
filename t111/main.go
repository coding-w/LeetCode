package main

func main() {

}

// https://leetcode.cn/problems/minimum-depth-of-binary-tree/
// tips 时间复杂度偏高
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right != nil {
		return min(minDepth(root.Left), minDepth(root.Right)) + 1
	} else if root.Left != nil || root.Right != nil {
		return max(minDepth(root.Left), minDepth(root.Right)) + 1
	} else {
		return 1
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
