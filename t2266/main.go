package main

func main() {

}

const mod = 1_000_000_007
const mx = 100_001

var f = [mx]int{1, 1, 2, 3}
var g = f

// https://leetcode.cn/problems/count-number-of-texts/?envType=daily-question&envId=2025-01-19
func init() {
	for i := 4; i < mx; i++ {
		f[i] = (f[i-1] + f[i-2] + f[i-3]) % mod
		g[i] = (g[i-1] + g[i-2] + g[i-3] + g[i-4]) % mod
	}
}

func countTexts(s string) int {
	res, cnt := 1, 0
	for i, c := range s {
		cnt++
		if i == len(s)-1 || byte(c) != s[i+1] {
			if c != '7' && c != '9' {
				res = res * f[cnt] % mod
			} else {
				res = res * g[cnt] % mod
			}
			cnt = 0
		}
	}
	return res
}
