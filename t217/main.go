package main

func main() {

}

// https://leetcode.cn/problems/contains-duplicate/description/
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = true
	}
	return false
}
