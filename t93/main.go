package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}

// https://leetcode.cn/problems/restore-ip-addresses/description/
// tips: 回溯
// 1 <= s.length <= 20
func restoreIpAddresses(s string) []string {
	n := len(s)
	// 长度小于或者大于16都是错误
	if n < 4 || n > 16 {
		return []string{}
	}
	res := make([]string, 0)
	tmp := make([]string, 0)
	// 回溯
	var dfs func(start int)
	// 判断是否合法
	var isValid func(s string) bool
	dfs = func(start int) {
		// 有四段，切已经用完
		if len(tmp) == 4 && start == n {
			res = append(res, strings.Join(tmp, "."))
			return
		}
		i := 1
		for ; i < 4; i++ {
			if start+i > n {
				continue
			}
			item := s[start : start+i]
			if !isValid(item) {
				continue
			}
			tmp = append(tmp, item)
			dfs(start + i)
			tmp = tmp[:len(tmp)-1]
		}
	}
	isValid = func(s string) bool {
		if len(s) > 1 && s[0] == '0' {
			return false
		}
		num, _ := strconv.Atoi(s)
		if num > 255 {
			return false
		}
		return true
	}
	dfs(0)
	return res
}
