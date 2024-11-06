package main

func main() {

}

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node/description/
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
	}
	if root.Next != nil {
		temp := root.Next
		for temp.Next != nil && temp.Next.Left == nil && temp.Next.Right == nil {
			temp = temp.Next
		}
		if temp.Next.Left != nil && root.Right != nil {
			root.Right.Next = temp.Next.Left
		} else if temp.Next.Right != nil && root.Right != nil {
			root.Right.Next = temp.Next.Right
		} else if temp.Next.Left != nil && root.Left != nil {
			root.Left.Next = temp.Next.Left
		} else if temp.Next.Right != nil && root.Left != nil {
			root.Left.Next = temp.Next.Right
		}
	}
	connect(root.Left)
	connect(root.Right)
	return root
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
