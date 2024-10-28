package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDecodings("226"))
}

// https://leetcode.cn/problems/decode-ways/description/
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		twoDigit := s[i-2 : i]
		if num, err := strconv.Atoi(twoDigit); err == nil && num >= 10 && num <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
