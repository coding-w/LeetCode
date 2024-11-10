package main

import "fmt"

func main() {
	s := "aab"
	fmt.Println(partition(s))

}

// https://leetcode.cn/problems/palindrome-partitioning/
func partition(s string) [][]string {
	l := len(s)
	if l <= 1 {
		return [][]string{{s}}
	}
	res := make([][]string, 0)
	tmp := make([]string, 0)
	var dfs func(int)
	dfs = func(i int) {
		if i == l {
			res = append(res, append([]string(nil), tmp...))
		}
		for j := i; j < l; j++ {
			t := s[i : j+1]
			if valid(t) {
				tmp = append(tmp, t)
				dfs(j + 1)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	dfs(0)
	return res
}
func valid(s string) bool {
	if len(s) <= 1 {
		return true
	}
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
