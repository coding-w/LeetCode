package main

import "fmt"

func main() {
	fmt.Println(1 & 5 & 3)
	fmt.Println(62 & 12 & 24 & 14 & 71)

}

// https://leetcode.cn/problems/largest-combination-with-bitwise-and-greater-than-zero/description/?envType=daily-question&envId=2025-01-12
// 每日一题
// 找到最大二进制的每一位都是1的最大值
// 十的七次方，大于2的23次方，小于2的24方
func largestCombination(candidates []int) int {
	res := 0
	for i := 1; i <= 24; i++ {
		cur := 0
		for _, candidate := range candidates {
			cur += (candidate >> i) & 1
		}
		if cur > res {
			res = cur
		}
	}
	return res
}
