package main

import (
	"fmt"
)

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
func (pq *PriorityQueue) Pop() (int, error) {
	if pq.size == 0 {
		return 0, fmt.Errorf("priority queue is empty")
	}
	// 堆顶元素
	top := pq.data[0]
	// 将堆尾元素放到堆顶
	pq.data[0] = pq.data[pq.size-1]
	// 移除堆尾
	pq.data = pq.data[:pq.size-1]
	pq.size--
	pq.siftDown(0) // 调整堆
	return top, nil
}

// Peek 获取堆顶元素
func (pq *PriorityQueue) Peek() (int, error) {
	if pq.size == 0 {
		return 0, fmt.Errorf("priority queue is empty")
	}
	return pq.data[0], nil
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

func main() {
	// 创建大顶堆
	maxHeap := NewPriorityQueue(func(a, b int) bool {
		// 大顶堆：父节点大于子节点
		return a > b
	})
	maxHeap.Push(3)
	maxHeap.Push(4)
	maxHeap.Push(5)
	maxHeap.Push(6)
	maxHeap.Push(1)
	maxHeap.Push(7)
	maxHeap.Push(8)
	peek, _ := maxHeap.Peek()
	fmt.Println("大顶堆 Peek:", peek)
	fmt.Println("大顶堆 Pop:")
	for maxHeap.size > 0 {
		val, _ := maxHeap.Pop()
		fmt.Print(val, " ")
	}
	fmt.Println()

	// 创建小顶堆
	minHeap := NewPriorityQueue(func(a, b int) bool {
		// 小顶堆：父节点小于子节点
		return a < b
	})
	minHeap.Push(3)
	minHeap.Push(4)
	minHeap.Push(5)
	minHeap.Push(6)
	minHeap.Push(1)
	minHeap.Push(7)
	minHeap.Push(8)
	peek, _ = minHeap.Peek()
	fmt.Println("小顶堆 Peek:", peek)
	fmt.Println("小顶堆 Pop:")
	for minHeap.size > 0 {
		val, _ := minHeap.Pop()
		fmt.Print(val, " ")
	}
}
