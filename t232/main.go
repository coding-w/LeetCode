package main

func main() {

}

// https://leetcode.cn/problems/implement-queue-using-stacks/
type MyQueue struct {
	list []int
}

func Constructor() MyQueue {
	return MyQueue{
		list: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
}

func (this *MyQueue) Pop() int {
	res := this.list[0]
	this.list = this.list[1:]
	return res
}

func (this *MyQueue) Peek() int {
	return this.list[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.list) == 0
}
