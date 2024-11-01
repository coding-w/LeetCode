package main

func main() {

}

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		if res < i {
			res = i
		}
		dfs(node.Left, i+1)
		dfs(node.Right, i+1)
	}
	dfs(root, 1)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
