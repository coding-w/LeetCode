package main

import "fmt"

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
}

// https://leetcode.cn/problems/isomorphic-strings/description/
func isIsomorphic(s string, t string) bool {
	// 建立两个映射表
	m1 := make(map[byte]byte) // s -> t 的映射
	m2 := make(map[byte]byte) // t -> s 的映射

	for i := 0; i < len(s); i++ {
		// 检查 s[i] -> t[i] 是否符合
		if v, ok := m1[s[i]]; ok && v != t[i] {
			return false
		}
		// 检查 t[i] -> s[i] 是否符合
		if v, ok := m2[t[i]]; ok && v != s[i] {
			return false
		}
		// 更新映射关系
		m1[s[i]] = t[i]
		m2[t[i]] = s[i]
	}
	return true
}
