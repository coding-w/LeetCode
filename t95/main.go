package main

import "fmt"

func main() {
	trees := generateTrees(3)
	for _, v := range trees {
		fmt.Println(printA(v))
	}
}

// https://leetcode.cn/problems/unique-binary-search-trees-ii/description/
// todo 未完成
func generateTrees(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 1}}
	}
	res := make([]*TreeNode, 0)
	head := &TreeNode{Val: n}
	var dfs func(int, *TreeNode)
	dfs = func(i int, node *TreeNode) {
		if i == 0 {
			res = append(res, head)
			fmt.Println(head)
			return
		}
		item := &TreeNode{Val: i}
		node.Left = item
		dfs(i-1, node.Left)
		if i-1 > 1 {
			item = &TreeNode{Val: i}
			item.Right = item
			dfs(i-1, node.Left)
		}
		node.Left = nil
		node.Right = item
		dfs(i-1, node.Right)
		node.Right = nil
	}
	dfs(n-1, head)

	return res
}
func printA(root *TreeNode) []int {
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
