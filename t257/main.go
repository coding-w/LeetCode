package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	fmt.Println(binaryTreePaths(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-paths/description/
func binaryTreePaths(root *TreeNode) (res []string) {
	if root == nil {
		return
	}
	left := binaryTreePaths(root.Left)
	right := binaryTreePaths(root.Right)
	itoa := strconv.Itoa(root.Val)
	for _, item := range left {
		res = append(res, itoa+"->"+item)
	}
	for _, item := range right {
		res = append(res, itoa+"->"+item)
	}
	if len(res) == 0 {
		res = append(res, itoa)
	}
	return
}
