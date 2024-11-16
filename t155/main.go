package main

import "math"

func main() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

}

// MinStack
// https://leetcode.cn/problems/min-stack/
type MinStack struct {
	nums   []int
	minNum int
}

func Constructor() MinStack {
	return MinStack{
		nums:   []int{},
		minNum: math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	this.minNum = min(this.minNum, val)
	this.nums = append(this.nums, val)
}

func (this *MinStack) Pop() {
	this.minNum = math.MaxInt32
	for i := 0; i < len(this.nums)-1; i++ {
		this.minNum = min(this.minNum, this.nums[i])
	}
	this.nums = this.nums[:len(this.nums)-1]
}

func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {
	return this.minNum
}
