package main

func main() {

}

func maxPathSum(root *TreeNode) int {
	maxSum := -1001
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sumLeft := dfs(node.Left)
		sumRight := dfs(node.Right)
		maxSum = max(maxSum, sumLeft+sumRight+node.Val)
		return max(max(sumLeft, sumRight)+node.Val, 0)
	}
	dfs(root)
	return maxSum
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
