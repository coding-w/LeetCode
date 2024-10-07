package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply("123", "523"))
}

// https://leetcode.cn/problems/multiply-strings/description/
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	// n位数 * n位数 为 2 n位数
	var res = make([]int, len(num1)+len(num2))
	n2 := len(num2)
	n1 := len(num1)
	for i := 0; i < n2; i++ {
		// 从右往左
		num := int(num2[n2-i-1] - '0')
		for j := 0; j < n1; j++ {
			// 计算结果
			res[i+j] += int(num1[n1-j-1]-'0') * num
			res[i+j+1] += res[i+j] / 10
			res[i+j] %= 10
		}
	}
	str := strings.Builder{}
	for i := len(res) - 1; i >= 0; i-- {
		if i == len(res)-1 && res[i] == 0 {
			continue
		}
		str.WriteByte(byte(res[i] + '0'))
	}
	return str.String()
}
