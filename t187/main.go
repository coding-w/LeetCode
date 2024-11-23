package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}

// https://leetcode.cn/problems/repeated-dna-sequences/
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	m := make(map[string]int)
	res := make([]string, 0)
	for i := 0; i <= len(s)-10; i++ {
		m[s[i:i+10]]++
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
