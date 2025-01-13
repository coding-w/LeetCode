package main

import (
	"fmt"
	"math"
)

const target = 2

func main() {
	fmt.Println(guessNumber(math.MaxInt32))
}

// https://leetcode.cn/problems/guess-number-higher-or-lower/description/
func guessNumber(n int) int {
	left := 1
	mid := (left + n) / 2
	res := guess(mid)
	count := 1
	for res != 0 {
		count++
		if res > 0 {
			left = mid + 1
			mid = (left + n) / 2
		} else {
			n = mid - 1
			mid = (left + n) / 2
		}
		res = guess(mid)
	}
	fmt.Println(count)
	return mid
}

func guess(num int) int {
	if num > target {
		return -1
	} else if num < target {
		return 1
	}
	return 0
}
