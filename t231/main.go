package main

func main() {

}

// https://leetcode.cn/problems/power-of-two/
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
