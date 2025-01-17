package main

import "math"

func main() {

}

func minimumSubarrayLength(nums []int, k int) int {
	bits := make([]int, 32)
	left, cur, n, res := 0, 0, len(nums), math.MaxInt32
	var updateBits func(int)
	updateBits = func(num int) {
		for i := 0; i < 32; i++ {
			if (num & (1 << i)) != 0 {
				bits[i]++
			}
		}
	}
	var delBits func(int)
	delBits = func(num int) {
		for i := 0; i < 32; i++ {
			if (num & (1 << i)) != 0 {
				bits[i]--
			}
		}
	}
	var check func() bool
	check = func() bool {
		tmp := 0
		for i := 0; i < 32; i++ {
			if bits[i] != 0 {
				tmp += 1 << i
			}
		}
		return tmp >= k
	}
	for cur < n {
		if nums[cur] >= k {
			return 1
		}
		updateBits(nums[cur])
		for left < n && check() {
			res = min(res, cur-left+1)
			delBits(nums[left])
			left++
		}
		cur++
	}
	if res == math.MaxInt32 {
		return -1
	}
	return res
}
