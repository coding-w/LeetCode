package main

import "fmt"

func main() {
	tree := &TreeNode{Left: nil, Right: &TreeNode{Left: &TreeNode{Left: nil, Right: nil, Val: 3}, Right: nil, Val: 2}, Val: 1}
	fmt.Println(inorderTraversal(tree))
}

// https://leetcode.cn/problems/binary-tree-inorder-traversal/
// tips: 中序遍历，左 中 右
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
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
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
