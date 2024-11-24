package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

// https://leetcode.cn/problems/count-primes/
func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	isPrime := make([]bool, n)
	for i := 2; i*i < n; i++ {
		if !isPrime[i] {
			// 设置 i 的所有倍数都是 非质数
			for j := i * i; j < n; j += i {
				isPrime[j] = true
			}
		}
	}
	count := 0
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			count++
		}
	}
	return count
}
