package main

func main() {

}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {

}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

// NestedIterator
// https://leetcode.cn/problems/flatten-nested-list-iterator/
// 实现一个迭代嵌套集合
type NestedIterator struct {
	index int
	nums  []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	c := &NestedIterator{
		index: 0,
		nums:  make([]int, 0),
	}
	c.dfs(nestedList)
	c.index = len(c.nums)
	return c

}

func (this *NestedIterator) dfs(nestedList []*NestedInteger) {
	for _, n := range nestedList {
		if n.IsInteger() {
			this.nums = append(this.nums, n.GetInteger())
		} else {
			this.dfs(n.GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	num := this.nums[this.index]
	this.index = this.index + 1
	return num
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.nums)
}
