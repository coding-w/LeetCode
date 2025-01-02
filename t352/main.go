package main

import "fmt"

func main() {
	c := Constructor()
	c.AddNum(1)
	fmt.Println(c.GetIntervals())
	c.AddNum(3)
	fmt.Println(c.GetIntervals())
	c.AddNum(7)
	fmt.Println(c.GetIntervals())
	c.AddNum(2)
	fmt.Println(c.GetIntervals())
	c.AddNum(6)
	fmt.Println(c.GetIntervals())
}

// SummaryRanges
// https://leetcode.cn/problems/data-stream-as-disjoint-intervals/
// 暴力破解！！
type SummaryRanges struct {
	nums []int
}

func Constructor() SummaryRanges {
	return SummaryRanges{
		nums: []int{},
	}
}

func (this *SummaryRanges) AddNum(value int) {
	this.nums = append(this.nums, value)
	if len(this.nums) > 1 {
		i := len(this.nums) - 1
		for i > 0 {
			if this.nums[i] == this.nums[i-1] {
				this.nums = append(this.nums[:i], this.nums[i+1:]...)
				return
			}
			if this.nums[i] < this.nums[i-1] {
				this.nums[i], this.nums[i-1] = this.nums[i-1], this.nums[i]
				i--
			} else {
				return
			}
		}
	}
}

func (this *SummaryRanges) GetIntervals() (res [][]int) {
	if len(this.nums) == 0 {
		return
	}
	k := 0
	i := 1
	for ; i < len(this.nums); i++ {
		if this.nums[i-1]+1 != this.nums[i] {
			res = append(res, []int{this.nums[k], this.nums[i-1]})
			k = i
		}
	}
	res = append(res, []int{this.nums[k], this.nums[i-1]})
	return res
}
