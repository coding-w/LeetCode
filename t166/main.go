package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(fractionToDecimal(4, 333))
}

// https://leetcode.cn/problems/fraction-to-recurring-decimal/
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	if numerator%denominator == 0 {
		return strconv.Itoa(numerator / denominator)
	}

	sb := strings.Builder{}
	if numerator > 0 && denominator < 0 || numerator < 0 && denominator > 0 {
		sb.WriteString("-")
	}

	// 处理符号
	a := numerator
	b := denominator
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	sb.WriteString(strconv.Itoa(a / b))
	sb.WriteString(".")

	a %= b
	// 记录余数及其对应的位置
	m := make(map[int]int)
	m[a] = sb.Len()

	for a != 0 {
		a *= 10
		sb.WriteString(strconv.Itoa(a / b))
		a %= b

		// 检查是否有循环
		if _, ok := m[a]; ok {
			break
		}

		m[a] = sb.Len()
	}

	if a != 0 {
		sb.WriteString(")")
		result := sb.String()
		result = result[:m[a]] + "(" + result[m[a]:]
		return result
	}

	return sb.String()
}
