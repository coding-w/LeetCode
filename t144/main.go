package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-preorder-traversal/description/
func preorderTraversal(root *TreeNode) (res []int) {
	//var dfs func(*TreeNode, int)
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return
}
