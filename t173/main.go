package main

func main() {

}

type BSTIterator struct {
	res []int
}

func Constructor(root *TreeNode) BSTIterator {
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
	return BSTIterator{res}
}

func (this *BSTIterator) Next() int {
	num := this.res[0]
	this.res = this.res[1:]
	return num
}

func (this *BSTIterator) HasNext() bool {
	return len(this.res) > 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
