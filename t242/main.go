package main

func main() {
}

// https://leetcode.cn/problems/valid-anagram/
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hmap := make(map[rune]int)
	for _, ch := range s {
		hmap[ch]++
	}
	for _, ch := range t {
		hmap[ch]--
	}
	for _, v := range hmap {
		if v != 0 {
			return false
		}
	}
	return true
}
