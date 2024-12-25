package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]string{"a", "aa", "aaa", "aaaa"}))
}

// https://leetcode.cn/problems/maximum-product-of-word-lengths/
// 暴力破解
func maxProduct(words []string) int {

	res := 0

	for i := 0; i < len(words)-1; i++ {
		word := words[i]
		h := make(map[uint8]bool)
		for j := 0; j < len(word); j++ {
			h[word[j]] = true
		}
		for j := i + 1; j < len(words); j++ {
			tmp := words[j]
			k := 0
			for k < len(tmp) {
				if h[tmp[k]] {
					break
				}
				k++
			}
			if k == len(tmp) {
				res = max(res, k*len(word))
			}
		}
	}

	return res
}
