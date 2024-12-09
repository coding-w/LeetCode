package main

import "fmt"

func main() {
	fmt.Println(constructTransformedArray([]int{3, -2, 1, 1}))
}

func constructTransformedArray(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		j := i + nums[i]
		for j < 0 {
			j += n
		}
		res[i] = nums[j%n]
	}
	return res
}
