package main

import "fmt"

func main() {
	fmt.Println(superPow(2, []int{1, 0}))
}

func superPow(a int, b []int) int {
	ans := 1
	for i := len(b) - 1; i >= 0; i-- {
		ans = ans * myPow(a, b[i]) % 1337
		a = myPow(a, 10)
	}
	return ans
}

func myPow(x int, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n&1 > 0 {
			res = res * x % 1337
		}
		x = x * x % 1337
	}
	return res
}
