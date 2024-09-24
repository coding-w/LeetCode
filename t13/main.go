package main

import "fmt"

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}

// https://leetcode.cn/problems/roman-to-integer/
func romanToInt(s string) int {
	rmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	for i := 0; i < len(s); i++ {
		res += rmap[s[i]]
		if i > 0 {
			if (s[i] == 'V' || s[i] == 'X') && s[i-1] == 'I' {
				res -= 2 * rmap[s[i-1]]
			}
			if (s[i] == 'L' || s[i] == 'C') && s[i-1] == 'X' {
				res -= 2 * rmap[s[i-1]]
			}
			if (s[i] == 'D' || s[i] == 'M') && s[i-1] == 'C' {
				res -= 2 * rmap[s[i-1]]
			}

		}
	}
	return res
}
