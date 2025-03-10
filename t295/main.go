package main

import (
	"fmt"
)

func main() {
	m := Constructor()
	m.AddNum(1)
	m.AddNum(2)
	fmt.Println(m.FindMedian())
	m.AddNum(3)
	fmt.Println(m.FindMedian())
}

// MedianFinder
// https://leetcode.cn/problems/find-median-from-data-stream/
type MedianFinder struct {
	left  *PriorityQueue
	right *PriorityQueue
}

func Constructor() MedianFinder {
	return MedianFinder{
		left: NewPriorityQueue(func(a, b int) bool {
			return a < b
		}),
		right: NewPriorityQueue(func(a, b int) bool {
			return a > b
		}),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.left.size == this.right.size {
		this.right.Push(num)
		this.left.Push(this.right.Pop())
	} else {
		this.left.Push(num)
		this.right.Push(this.left.Pop())
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.size == this.right.size {
		return float64(this.left.Peek()+this.right.Peek()) / 2
	} else {
		return float64(this.left.Peek())
	}
}

// PriorityQueue 定义优先队列结构
type PriorityQueue struct {
	data     []int               // 堆数据存储
	size     int                 // 当前堆大小
	lessFunc func(a, b int) bool // 比较函数，决定堆的顺序
}

// NewPriorityQueue 创建一个优先队列
func NewPriorityQueue(lessFunc func(a, b int) bool) *PriorityQueue {
	return &PriorityQueue{
		data:     []int{},
		size:     0,
		lessFunc: lessFunc,
	}
}

// Push 插入元素并调整堆（上滤）
func (pq *PriorityQueue) Push(val int) {
	pq.data = append(pq.data, val) // 将元素添加到堆末尾
	pq.size++
	pq.siftUp(pq.size - 1) // 调整堆
}

// Pop 删除堆顶元素并调整堆（下滤）
func (pq *PriorityQueue) Pop() int {
	if pq.size == 0 {
		return 0
	}
	// 堆顶元素
	top := pq.data[0]
	// 将堆尾元素放到堆顶
	pq.data[0] = pq.data[pq.size-1]
	// 移除堆尾
	pq.data = pq.data[:pq.size-1]
	pq.size--
	pq.siftDown(0) // 调整堆
	return top
}

// Peek 获取堆顶元素
func (pq *PriorityQueue) Peek() int {
	if pq.size == 0 {
		return 0
	}
	return pq.data[0]
}

// 上滤操作
func (pq *PriorityQueue) siftUp(index int) {
	for index > 0 {
		// 父节点索引
		parent := (index - 1) / 2
		// 如果当前节点与父节点满足堆的顺序，停止上滤
		if !pq.lessFunc(pq.data[index], pq.data[parent]) {
			break
		}
		// 否则交换父子节点并继续
		pq.data[index], pq.data[parent] = pq.data[parent], pq.data[index]
		index = parent
	}
}

// 下滤操作
func (pq *PriorityQueue) siftDown(index int) {
	for index*2+1 < pq.size {
		left := index*2 + 1  // 左子节点索引
		right := index*2 + 2 // 右子节点索引
		smallest := left     // 假设左子节点是较小的

		// 如果右子节点存在且更符合堆顺序规则
		if right < pq.size && pq.lessFunc(pq.data[right], pq.data[left]) {
			smallest = right
		}

		// 如果当前节点与子节点满足堆的顺序，停止下滤
		if !pq.lessFunc(pq.data[smallest], pq.data[index]) {
			break
		}

		// 否则交换当前节点与较小的子节点
		pq.data[index], pq.data[smallest] = pq.data[smallest], pq.data[index]
		index = smallest
	}
}
