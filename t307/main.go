package main

func main() {
	c := Constructor([]int{1, 3, 5})
	println(c.SumRange(0, 2))
	c.Update(1, 2)
	println(c.SumRange(0, 2))
}

// NumArray https://leetcode.cn/problems/range-sum-query-mutable/
type NumArray struct {
	sums []int
	nums []int
}

func Constructor(nums []int) NumArray {
	tmp := make([]int, len(nums)+1)
	sum := 0
	for i := 1; i < len(tmp); i++ {
		sum += nums[i-1]
		tmp[i] = sum
	}
	return NumArray{sums: tmp, nums: nums}
}

func (this *NumArray) Update(index int, val int) {
	diff := val - this.nums[index]
	this.nums[index] = val
	for i := index + 1; i < len(this.sums); i++ {
		this.sums[i] += diff
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(index,val);
 * param_2 := obj.SumRange(left,right);
 */
