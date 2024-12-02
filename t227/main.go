package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculate("3+2*3"))

}

// https://leetcode.cn/problems/basic-calculator-ii/description/
// tips 中序转后续
func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	nums := make([]int, 0)
	num, index := getNum(s, 0)
	i := 0
	nums = append(nums, num)
	for index < len(s) {
		if s[index] > '9' || s[index] < '0' {
			num, i = getNum(s, index+1)
			switch s[index] {
			case '+':
				nums = append(nums, num)
			case '-':
				nums = append(nums, -num)
			case '*':
				num = nums[len(nums)-1] * num
				nums[len(nums)-1] = num
			case '/':
				num = nums[len(nums)-1] / num
				nums[len(nums)-1] = num
			}
			index = i
		}

	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum
}

func getNum(s string, index int) (int, int) {
	start := index
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		index++
	}
	num, _ := strconv.Atoi(s[start:index])
	return num, index
}
