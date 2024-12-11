package main

func main() {

}

// https://leetcode.cn/problems/first-bad-version/
func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func isBadVersion(version int) bool {
	return false
}
