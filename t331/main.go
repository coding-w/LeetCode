package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isValidSerialization("9,3,4,#,#,1,#,#,2,#,6,#,#"))
}

// https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/
// 9,3,4,#,#,1,#,#,2,#,6,#,#
// 每一个叶子节点都有两个 '#',遇到两个 # 消除三个
func isValidSerialization(preorder string) bool {
	str := strings.Split(preorder, ",")
	length := len(str)
	i := 0
	for {
		if i >= length-1 {
			break
		}
		if str[0] == "#" && len(str) > 1 {
			return false
		}
		if str[i] == "#" && str[i+1] == "#" && str[i-1] != "#" {
			str = append(str[:i], str[i+2:]...)
			str[i-1] = "#"
			i = 0
			length -= 2
			continue
		}
		i++
	}
	return len(str) == 1 && str[0] == "#"
}
