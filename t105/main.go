package main

func main() {

}

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
func buildTree(preorder []int, inorder []int) *TreeNode {
	m := make(map[int]int, len(inorder))
	for i, v := range inorder {
		m[v] = i
	}
	var dfs func(int, int) *TreeNode
	i := 0
	dfs = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		root := preorder[i]
		i++
		return &TreeNode{Val: root, Left: dfs(left, m[root]-1), Right: dfs(m[root]+1, right)}
	}
	return dfs(0, len(preorder)-1)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
