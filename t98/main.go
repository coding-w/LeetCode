package main

import "fmt"

func main() {
	tree := &TreeNode{Left: nil, Right: &TreeNode{Left: &TreeNode{Left: nil, Right: nil, Val: 3}, Right: nil, Val: 2}, Val: 1}
	fmt.Println(isValidBST(tree))
}

// https://leetcode.cn/problems/validate-binary-search-tree/description/
// tips: 中序遍历，左 中 右
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	for i := range res {
		if i == 0 {
			continue
		}
		if res[i] <= res[i-1] {
			return false
		}
	}
	return true
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
