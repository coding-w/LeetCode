package main

func main() {

}

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
func sortedArrayToBST(nums []int) *TreeNode {
	l := len(nums)
	var dfs func(left int, right int) *TreeNode
	dfs = func(left int, right int) *TreeNode {
		if left >= right {
			return nil
		}
		mid := left + (right-left)/2
		tree := &TreeNode{Val: nums[mid]}
		tree.Left = dfs(left, mid)
		tree.Right = dfs(mid+1, right)
		return tree
	}
	return dfs(0, l)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
