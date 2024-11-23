package main

import "fmt"

func main() {
	fmt.Println(countValidSelections([]int{2, 3, 4, 0, 4, 1, 0}))
}

func countValidSelections(nums []int) int {

	res := 0
	var dfs func([]int, int, int)
	dfs = func(nums []int, i int, j int) {
		if i >= len(nums) || i < 0 {
			for _, num := range nums {
				if num != 0 {
					return
				}
			}
			res++
			return
		}
		if nums[i] == 0 {
			dfs(nums, i+j, j)
		}
		if nums[i] > 0 {
			nums[i]--
			dfs(nums, i-j, -j)
			nums[i]++
		}
	}
	for i, v := range nums {
		if v == 0 {
			dfs(nums, i+1, 1)
			dfs(nums, i-1, -1)
		}
	}

	return res

}
