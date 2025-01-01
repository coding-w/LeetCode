package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
}

// https://leetcode.cn/problems/convert-date-to-binary/?envType=daily-question&envId=2025-01-01
// 每日一题
func convertDateToBinary(date string) string {
	splits := strings.Split(date, "-")
	res := make([]string, 3)
	for i, s := range splits {
		atoi, _ := strconv.Atoi(s)
		res[i] = fmt.Sprintf("%b", atoi)
	}
	return strings.Join(res, "-")
}
