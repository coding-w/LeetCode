package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Codec 类
type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	var sb []string

	// 前序遍历序列化
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			sb = append(sb, "null") // 空节点
			return
		}
		sb = append(sb, strconv.Itoa(node.Val)) // 非空节点
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return strings.Join(sb, ",") // 生成序列化字符串
}

// 反序列化函数
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	// 将字符串分割为数组
	nodes := strings.Split(data, ",")
	index := 0

	// 递归构建二叉树
	var buildTree func() *TreeNode
	buildTree = func() *TreeNode {
		if nodes[index] == "null" {
			index++
			return nil
		}

		val, _ := strconv.Atoi(nodes[index])
		index++
		node := &TreeNode{Val: val}
		node.Left = buildTree()
		node.Right = buildTree()
		return node
	}

	return buildTree()
}

// 测试代码
func main() {
	// 创建测试树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	// 序列化和反序列化
	codec := Constructor()
	serialized := codec.serialize(root)
	fmt.Println("序列化:", serialized)

	deserialized := codec.deserialize(serialized)
	fmt.Println("反序列化后的树（中序遍历）:")
	printInOrder(deserialized)
}

// 辅助函数：中序遍历打印树
func printInOrder(root *TreeNode) {
	if root == nil {
		return
	}
	printInOrder(root.Left)
	fmt.Print(root.Val, " ")
	printInOrder(root.Right)
}
