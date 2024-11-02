package main

func main() {

}

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii/
// tips:遍历记住层数
func levelOrderBottom(root *TreeNode) [][]int {
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
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
