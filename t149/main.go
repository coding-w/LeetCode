package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(maxPoints([][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}))
}

// https://leetcode.cn/problems/max-points-on-a-line/
// tips: 斜率
func maxPoints(points [][]int) int {
	mx := 0
	for i := 0; i < len(points); i++ {
		m := make(map[string]int)
		for j := i + 1; j < len(points); j++ {
			y := points[i][1] - points[j][1]
			x := points[i][0] - points[j][0]
			z := gcd(x, y)
			key := strconv.Itoa(y/z) + "/" + strconv.Itoa(x/z)
			m[key]++
			mx = max(mx, m[key])
		}
	}

	return mx
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
