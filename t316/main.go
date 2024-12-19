package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
}

// https://leetcode.cn/problems/remove-duplicate-letters/description/
func removeDuplicateLetters(s string) string {

	// 每一个字符最后出现的位置
	hmap := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		hmap[s[i]] = i
	}

	// 定义栈
	stack := make([]uint8, 0)
	// 定义某个字符是否保存在栈中
	isExist := make(map[uint8]bool)

	for i := 0; i < len(s); i++ {
		if !isExist[s[i]] {
			for len(stack) > 0 && stack[len(stack)-1] > s[i] && hmap[stack[len(stack)-1]] > i {
				v := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				isExist[v] = false
			}
			isExist[s[i]] = true
			stack = append(stack, s[i])
		}
	}
	res := strings.Builder{}
	for i := 0; i < len(stack); i++ {
		res.WriteString(string(stack[i]))

	}
	return res.String()

}
