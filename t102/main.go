package main

func main() {

}

// https://leetcode.cn/problems/same-tree/
// tips:遍历记住层数
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)

	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, i int) {
		if node == nil {
			return
		}
		if len(res) < i {
			res = append(res, []int{})
		}
		res[i-1] = append(res[i-1], node.Val)
		dfs(node.Left, i+1)
		dfs(node.Right, i+1)
	}
	dfs(root, 1)
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
