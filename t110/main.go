package main

import "math"

func main() {

}

// https://leetcode.cn/problems/balanced-binary-tree/
func isBalanced(root *TreeNode) bool {
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := height(root.Left)
		rightHeight := height(root.Right)
		if leftHeight >= 0 && rightHeight >= 0 && math.Abs(float64(leftHeight)-float64(rightHeight)) <= 1 {
			return max(leftHeight, rightHeight) + 1
		} else {
			return -1
		}
	}
	return height(root) >= 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
