package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(compareVersion("1.2.3", "1.2.3.0.0"))
}

// https://leetcode.cn/problems/compare-version-numbers/
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	l := max(len(v1), len(v2))
	for i := 0; i < l; i++ {
		r1, r2 := 0, 0
		if i < len(v1) {
			r1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			r2, _ = strconv.Atoi(v2[i])
		}
		if r1 > r2 {
			return 1
		}
		if r1 < r2 {
			return -1
		}
	}
	return 0
}
