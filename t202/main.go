package main

func main() {

}

// https://leetcode.cn/problems/happy-number/
func isHappy(n int) bool {
	for i := 0; i < 10; i++ {
		sum := 0
		for n > 0 {
			k := n % 10
			n /= 10
			sum += k * k
		}
		if sum == 1 {
			return true
		}
		n = sum
	}
	return n == 1
}
