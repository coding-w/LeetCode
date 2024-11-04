package main

func main() {

}

// https://leetcode.cn/problems/path-sum/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		if 0 == targetSum {
			return true
		}
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
