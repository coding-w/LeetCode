package main

func main() {
	node1 := &Node{Val: 12}
	node1.Next = &Node{Val: 3, Next: nil, Random: node1}
	copyRandomList(node1)
}

// https://leetcode.cn/problems/copy-list-with-random-pointer/
func copyRandomList(head *Node) *Node {
	res := new(Node)
	p, q := res, head
	m := make(map[*Node]*Node)
	for q != nil {
		p.Next = &Node{Val: q.Val}
		p = p.Next
		m[q] = p
		q = q.Next
	}
	p, q = res.Next, head
	for q != nil {
		if q.Random != nil {
			p.Random = m[q.Random]
		}
		q = q.Next
		p = p.Next
	}
	return res.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
