package main

import "fmt"

func main() {
	root := &TreeNode{Val: 5,
		Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}},
		Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}},
	}
	fmt.Println(pathSum(root, 22))
}

// https://leetcode.cn/problems/path-sum/
func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	var dfs func(*TreeNode, int, []int)
	dfs = func(node *TreeNode, target int, tmp []int) {
		if node == nil {
			return
		}
		target -= node.Val
		tmp = append(tmp, node.Val)
		if node.Left == nil && node.Right == nil && 0 == target {
			res = append(res, append([]int{}, tmp...))
			return
		} else {
			dfs(node.Left, target, tmp)
			dfs(node.Right, target, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}
	dfs(root, targetSum, []int{})
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
