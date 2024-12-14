package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(getHint("1234", "4321"))
}

// https://leetcode.cn/problems/bulls-and-cows/
func getHint(secret string, guess string) string {
	// 记录 A 和 B 的数量
	countA, countB := 0, 0
	// 数字
	nums := make([]int, 10)
	for i := 0; i < len(guess); i++ {
		if guess[i] == secret[i] {
			countA++
			continue
		}
		// 数字
		a := guess[i] - '0'
		b := secret[i] - '0'
		// 数字是否都存在
		if nums[a] > 0 {
			countB++
		}
		if nums[b] < 0 {
			countB++
		}
		nums[a]--
		nums[b]++
	}
	return strconv.Itoa(countA) + "A" + strconv.Itoa(countB) + "B"

}
