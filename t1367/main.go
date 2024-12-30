package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://leetcode.cn/problems/linked-list-in-binary-tree/?envType=daily-question&envId=2024-12-30
// 每日一题
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	return dfs(root, head) || isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(node *TreeNode, head *ListNode) bool {
	// 链表已经全部匹配完，匹配成功
	if head == nil {
		return true
	}
	// 二叉树访问到了空节点，匹配失败
	if node == nil {
		return false
	}
	// 当前匹配的二叉树上节点的值与链表节点的值不相等，匹配失败
	if node.Val != head.Val {
		return false
	}
	return dfs(node.Left, head.Next) || dfs(node.Right, head.Next)
}
