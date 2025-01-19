package main

import (
	"fmt"
	"sort"
)

func minCost(arr []int, brr []int, k int64) int64 {
	ans := int64(0)
	for i := 0; i < len(arr); i++ {
		ans += abs(int64(arr[i]) - int64(brr[i]))
	}
	sort.Ints(arr)
	sort.Ints(brr)

	res := k

	for i := 0; i < len(arr); i++ {
		res += abs(int64(arr[i]) - int64(brr[i]))
	}

	return min(res, ans)
}

func abs(num int64) int64 {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// 测试例子
	arr := []int{-7, 9, 5}
	brr := []int{7, -2, -5}
	//k := 2

	result := minCost(arr, brr, 2)
	fmt.Println(result) // 输出 13
}
