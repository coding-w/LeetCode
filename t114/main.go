package main

import "fmt"

func main() {
	//root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	//root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	//root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 4}}}
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}}}
	flatten(root)
}

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/
// 花了好长时间....
// tips: 将左树，放在右树
func flatten(root *TreeNode) {
	// 空节点，或者只有一个节点
	if root == nil || root.Left == nil && root.Right == nil {
		return
	}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		// 空节点，或者只有一个节点 返回
		if root == nil || root.Left == nil && root.Right == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		if root.Left != nil && root.Left.Left == nil {
			// 保证左树没有左节点，将左树的右子节点 放在都放在 root 的右子节点
			tail := root.Left
			for tail.Right != nil {
				tail = tail.Right
			}
			tail.Right = root.Right
			root.Right = root.Left
			root.Left = nil
			return
		}
		if root.Right == nil && root.Left != nil {
			// 右树为空将左树直接放到左树下
			root.Right = root.Left
			root.Left = nil
			return
		}
	}
	dfs(root)
	//printP(root)
}
func printP(root *TreeNode) {
	if root == nil {
		return
	}
	res := make([]int, 0)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
