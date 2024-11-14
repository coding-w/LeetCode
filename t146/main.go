package main

import "fmt"

func main() {
	fmt.Println(hasIncreasingSubarrays([]int{-3, -19, -8, -16}, 2))
}
func hasIncreasingSubarrays(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	if k == 1 {
		return true
	}

	for i := 0; i < len(nums); i++ {
		if i+2*k <= len(nums) {
			j := i
			for ; j < i+k-1; j++ {
				if nums[j] >= nums[j+1] {
					break
				}
			}
			if j != i+k-1 || nums[j-1] >= nums[j] {
				continue
			}
			j = i + k
			for ; j < i+2*k-1; j++ {
				if nums[j] >= nums[j+1] {
					break
				}
			}
			if j == i+2*k-1 && nums[j-1] < nums[j] {
				return true
			}
		}
	}
	return false
}
