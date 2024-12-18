package main

func main() {

}

// https://leetcode.cn/problems/super-ugly-number/
func nthSuperUglyNumber(n int, primes []int) int {
	pq := NewPriorityQueue(func(a, b int64) bool {
		return a < b
	})
	pq.Push(1)
	hmap := make(map[int64]bool)
	hmap[1] = true
	var ugly int64 = 1
	for i := 0; i < n; i++ {
		ugly = pq.Pop()
		for _, v := range primes {
			newUgly := ugly * int64(v)
			if _, ok := hmap[newUgly]; !ok {
				hmap[newUgly] = true
				pq.Push(newUgly)
			}
		}
	}
	return int(ugly)
}

// PriorityQueue 定义优先队列结构
type PriorityQueue struct {
	data     []int64               // 堆数据存储
	size     int                   // 当前堆大小
	lessFunc func(a, b int64) bool // 比较函数，决定堆的顺序
}

// NewPriorityQueue 创建一个优先队列
func NewPriorityQueue(lessFunc func(a, b int64) bool) *PriorityQueue {
	return &PriorityQueue{
		data:     []int64{},
		size:     0,
		lessFunc: lessFunc,
	}
}

// Push 插入元素并调整堆（上滤）
func (pq *PriorityQueue) Push(val int64) {
	pq.data = append(pq.data, val) // 将元素添加到堆末尾
	pq.size++
	pq.siftUp(pq.size - 1) // 调整堆
}

// Pop 删除堆顶元素并调整堆（下滤）
func (pq *PriorityQueue) Pop() int64 {
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
func (pq *PriorityQueue) Peek() int64 {
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
