package main

func main() {

}

// NumArray https://leetcode.cn/problems/range-sum-query-immutable/
type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	tmp := make([]int, len(nums)+1)
	sum := 0
	for i := 1; i < len(tmp); i++ {
		sum += nums[i-1]
		tmp[i] = sum
	}
	return NumArray{tmp}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.nums[right+1] - this.nums[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
