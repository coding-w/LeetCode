package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/house-robber-iii/
// 超时 递归 + 记忆 完成 ...
func rob(root *TreeNode) int {
	h := make(map[*TreeNode]int)
	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if v, ok := h[node]; ok {
			return v
		}
		if node.Left == nil && node.Right == nil {
			h[node] = node.Val
			return node.Val
		}

		// 不选本节点
		r1 := dfs(node.Left) + dfs(node.Right)

		r2 := node.Val

		if node.Left != nil {
			r2 += dfs(node.Left.Left) + dfs(node.Left.Right)
		}
		if node.Right != nil {
			r2 += dfs(node.Right.Left) + dfs(node.Right.Right)
		}
		h[node] = max(r1, r2)
		return h[node]
	}
	return dfs(root)

}
