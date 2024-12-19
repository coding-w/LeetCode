package main

import (
	"fmt"
)

func main() {
	fmt.Println(2e4)
	fmt.Println(countSmaller([]int{6, 2}))
}

// https://leetcode.cn/problems/count-of-smaller-numbers-after-self/
func countSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// 初始化
	initArr := make([][]int, n)
	for i := 0; i < n; i++ {
		initArr[i] = make([]int, 2)
		initArr[i][0] = i
		initArr[i][1] = nums[i]
	}

	var mergeSort func(arr [][]int) [][]int
	var merge func(left, right [][]int) [][]int
	mergeSort = func(arr [][]int) [][]int {
		length := len(arr)
		if length <= 1 {
			return arr
		}
		mid := length / 2
		left := mergeSort(arr[:mid])
		right := mergeSort(arr[mid:])
		return merge(left, right)
	}
	merge = func(left, right [][]int) (tmp [][]int) {
		i, j := 0, 0
		for i < len(left) || j < len(right) {
			if j == len(right) || i < len(left) && left[i][1] <= right[j][1] {
				tmp = append(tmp, left[i])
				res[left[i][0]] += j
				i++
			} else {
				tmp = append(tmp, right[j])
				j++
			}
		}
		return
	}

	fmt.Println(mergeSort(initArr))
	//mergeSort(initArr)
	return res

}
