package main

import "fmt"

func main() {
	fmt.Println(isPerfectSquare(16))
}

// https://leetcode.cn/problems/valid-perfect-square/
func isPerfectSquare(num int) bool {
	left, right := 1, num
	for left <= right {
		mid := left + (right-left)/2
		t := num / mid
		if t == mid {
			if num%mid == 0 {
				return true
			} else {
				return false
			}
		} else if t < mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
