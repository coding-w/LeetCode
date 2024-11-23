package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/binary-tree-right-side-view/
func rightSideView(root *TreeNode) []int {
	tmp := make([][]int, 0)
	var dfs func(root *TreeNode, index int)
	dfs = func(root *TreeNode, index int) {
		if root == nil {
			return
		}
		if index == len(tmp) {
			item := make([]int, 0)
			item = append(item, root.Val)
			tmp = append(tmp, item)
		} else {
			tmp[index] = append(tmp[index], root.Val)
		}
		dfs(root.Left, index+1)
		dfs(root.Right, index+1)
	}
	dfs(root, 0)
	res := make([]int, 0)
	for _, v := range tmp {
		res = append(res, len(v)-1)
	}
	return res
}
