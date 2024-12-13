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
	arr  []int
	flag bool
	slow int
}

func Constructor() MedianFinder {
	return MedianFinder{
		arr:  []int{},
		flag: true,
		slow: -1,
	}
}

func (this *MedianFinder) AddNum(num int) {
	this.arr = append(this.arr, num)
	if this.flag {
		this.slow++
	}
	this.flag = !this.flag
}

func (this *MedianFinder) FindMedian() float64 {
	if this.flag {
		return float64(this.arr[this.slow]+this.arr[this.slow+1]) / 2
	}
	return float64(this.arr[this.slow])
}

type Heap []int
