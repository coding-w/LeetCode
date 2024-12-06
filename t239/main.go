package main

import "fmt"

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

type queueList struct {
	list []int
}

func (q *queueList) isEmpty() bool {
	return len(q.list) == 0
}

func (q *queueList) front() int {
	return q.list[0]
}

func (q *queueList) last() int {
	return q.list[len(q.list)-1]
}

func (q *queueList) pop(v int) {
	if q.front() == v {
		q.list = q.list[1:]
	}
}

func (q *queueList) push(v int) {
	for !q.isEmpty() && q.last() < v {
		q.list = q.list[:len(q.list)-1]
	}
	q.list = append(q.list, v)
}

// https://leetcode.cn/problems/sliding-window-maximum/
// 使用 单调队列完成
func maxSlidingWindow(nums []int, k int) []int {
	i, res := 0, make([]int, 0)
	list := &queueList{
		list: make([]int, 0),
	}
	for i < k {
		list.push(nums[i])
		i++
	}
	res = append(res, list.front())
	for i < len(nums) {
		list.pop(nums[i-k])
		list.push(nums[i])
		res = append(res, list.front())
		i++
	}
	return res
}
