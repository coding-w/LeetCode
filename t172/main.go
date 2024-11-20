package main

func main() {

}

// https://leetcode.cn/problems/factorial-trailing-zeroes/
// 求5的个数
func trailingZeroes(n int) int {
	res := 0
	for n >= 5 {
		res = res + n/5
		n /= 5
	}
	return res
}
