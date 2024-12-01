package main

func main() {

}

type MyStack struct {
	stack []int
	len   int
}

func Constructor() MyStack {
	return MyStack{make([]int, 0), 0}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
	this.len++

}

func (this *MyStack) Pop() int {
	res := this.stack[this.len-1]
	this.stack = this.stack[:this.len-1]
	this.len--
	return res

}

func (this *MyStack) Top() int {
	return this.stack[this.len-1]
}

func (this *MyStack) Empty() bool {
	return this.len == 0
}
