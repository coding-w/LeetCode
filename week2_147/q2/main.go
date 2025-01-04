package main

import (
	"container/heap"
	"fmt"
)

func main() {
	c := Constructor([][]int{{1, 101, 10}, {2, 102, 20}, {3, 103, 15}})
	c.Add(4, 104, 5)
	c.Edit(102, 8)
	fmt.Println(c.ExecTop())
	c.Rmv(101)
	c.Add(5, 101, 15)
	fmt.Println(c.ExecTop())
}

// Task 结构体表示单个任务
type Task struct {
	UserID   int // 任务所属用户
	TaskID   int // 任务 ID
	Priority int // 优先级
	Index    int // 用于堆中的索引（heap 内部使用）
}

// TaskManager 管理任务的核心结构
type TaskManager struct {
	taskMap   map[int]*Task   // 存储 TaskID -> Task 的映射，用于快速查找任务
	userTasks map[int][]*Task // 存储 UserID -> 任务列表
	taskHeap  *PriorityQueue  // 优先级队列，按任务优先级排序
}

// Constructor 构造函数，初始化任务管理器
func Constructor(tasks [][]int) TaskManager {
	manager := TaskManager{
		taskMap:   make(map[int]*Task),
		userTasks: make(map[int][]*Task),
		taskHeap:  &PriorityQueue{},
	}

	heap.Init(manager.taskHeap) // 初始化优先级队列

	// 初始化任务
	for _, t := range tasks {
		userID, taskID, priority := t[0], t[1], t[2]
		task := &Task{
			UserID:   userID,
			TaskID:   taskID,
			Priority: priority,
		}
		manager.taskMap[taskID] = task
		manager.userTasks[userID] = append(manager.userTasks[userID], task)
		heap.Push(manager.taskHeap, task)
	}
	return manager
}

// Add 向系统中添加一个新任务
func (m *TaskManager) Add(userID, taskID, priority int) {
	if _, exists := m.taskMap[taskID]; exists {
		return // 如果任务已经存在，不添加
	}
	task := &Task{
		UserID:   userID,
		TaskID:   taskID,
		Priority: priority,
	}
	m.taskMap[taskID] = task
	m.userTasks[userID] = append(m.userTasks[userID], task)
	heap.Push(m.taskHeap, task) // 将任务加入优先级队列
}

// Edit 修改任务的优先级
func (m *TaskManager) Edit(taskID, newPriority int) {
	if task, exists := m.taskMap[taskID]; exists {
		task.Priority = newPriority
		heap.Fix(m.taskHeap, task.Index) // 更新堆的顺序
	}
}

// Rmv 删除一个任务
func (m *TaskManager) Rmv(taskID int) {
	if task, exists := m.taskMap[taskID]; exists {
		heap.Remove(m.taskHeap, task.Index) // 从堆中删除任务
		delete(m.taskMap, taskID)           // 从任务映射中删除
	}
}

// ExecTop 执行所有任务中优先级最高的任务
func (m *TaskManager) ExecTop() int {
	if m.taskHeap.Len() == 0 {
		return -1 // 没有任务，返回 -1
	}
	task := heap.Pop(m.taskHeap).(*Task) // 弹出优先级最高的任务
	delete(m.taskMap, task.TaskID)       // 从任务映射中删除任务
	return task.UserID                   // 返回任务的所属用户
}

// PriorityQueue 优先级队列实现
type PriorityQueue []*Task

// Len 返回队列的长度
func (pq PriorityQueue) Len() int { return len(pq) }

// Less 比较两个任务的优先级（优先级高的任务优先）
func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].Priority == pq[j].Priority {
		return pq[i].TaskID > pq[j].TaskID // 如果优先级相同，按 TaskID 降序排序
	}
	return pq[i].Priority > pq[j].Priority
}

// Swap 交换两个任务
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

// Push 向队列中添加任务
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	task := x.(*Task)
	task.Index = n
	*pq = append(*pq, task)
}

// Pop 从队列中移除优先级最高的任务
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	task := old[n-1]
	old[n-1] = nil // 避免内存泄漏
	task.Index = -1
	*pq = old[0 : n-1]
	return task
}
