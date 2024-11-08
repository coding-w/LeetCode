package main

func main() {

}

// https://leetcode.cn/problems/sum-root-to-leaf-numbers/
func sumNumbers(root *TreeNode) int {
	var dfs func(root *TreeNode, sum int) int
	dfs = func(root *TreeNode, sum int) int {
		if root == nil {
			return 0
		}
		sum *= 10
		if root.Left == nil && root.Right == nil {
			return root.Val + sum
		}
		sum += root.Val
		return dfs(root.Left, sum) + dfs(root.Right, sum)
	}
	return dfs(root, 0)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
