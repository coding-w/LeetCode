package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "101"
	b := "1101"
	// 将二进制字符串转换为整数
	numA, _ := strconv.ParseInt(a, 2, 64) // 以2为基数，转换为64位整数
	numB, _ := strconv.ParseInt(b, 2, 64)
	// 将两个整数相加
	sum := numA + numB
	// 将相加结果转换回二进制字符串
	fmt.Println(strconv.FormatInt(sum, 2))
	fmt.Println(addBinary(a, b))
}

// https://leetcode.cn/problems/add-binary/description/
func addBinary(a string, b string) string {
	var sb strings.Builder
	i, j := len(a)-1, len(b)-1
	var tmp uint8 = 0
	for i >= 0 || j >= 0 {
		if i >= 0 {
			tmp += a[i] - '0'
		}
		if j >= 0 {
			tmp += b[j] - '0'
		}
		sb.WriteByte((tmp % 2) + '0')
		tmp /= 2
		i--
		j--
	}
	if tmp != 0 {
		sb.WriteByte((tmp % 2) + '0')
	}
	s := sb.String()
	runes := []rune(s)
	for i, j = 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
