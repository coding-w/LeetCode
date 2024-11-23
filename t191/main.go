package main

func main() {

}

// https://leetcode.cn/problems/number-of-1-bits/
func hammingWeight(n int) int {
	res := 0
	for n != 0 {
		res += n % 2
		n /= 2
	}
	return res

}
