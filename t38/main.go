package main

import "strconv"

func main() {
	print(countAndSay(4))
}

// https://leetcode.cn/problems/count-and-say/
func countAndSay(n int) string {

	if n == 1 {
		return "1"
	}

	return recursion(countAndSay(n - 1))

}

func recursion(s string) string {
	var res string
	var count int
	var last byte
	for i := 0; i < len(s); i++ {
		if i == 0 {
			count = 1
			last = s[i]
			continue
		}
		if s[i] == last {
			count++
			continue
		}
		res += strconv.Itoa(count) + string(last)
		count = 1
		last = s[i]

	}
	res += strconv.Itoa(count) + string(last)

	return res
}
