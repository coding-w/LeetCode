package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordPattern("abba", "dog dog dog dog"))
}

// https://leetcode.cn/problems/word-pattern/
func wordPattern(pattern string, s string) bool {
	split := strings.Split(s, " ")
	if len(pattern) != len(split) {
		return false
	}
	pm := make(map[uint8]string)
	sm := make(map[string]uint8)
	for i := 0; i < len(split); i++ {
		v1, ok1 := pm[pattern[i]]
		v2, ok2 := sm[split[i]]
		if (ok1 && v1 != split[i]) || (ok2 && v2 != pattern[i]) {
			return false
		}

		// 建立映射
		pm[pattern[i]] = split[i]
		sm[split[i]] = pattern[i]
	}
	return true
}
