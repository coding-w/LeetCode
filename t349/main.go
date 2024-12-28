package main

func main() {

}

// https://leetcode.cn/problems/intersection-of-two-arrays/
func intersection(nums1 []int, nums2 []int) []int {
	res := make([]int, 0)
	hash := make(map[int]bool)
	hash2 := make(map[int]bool)
	for _, num := range nums1 {
		hash[num] = true
	}
	for _, num := range nums2 {
		if hash[num] {
			hash2[num] = true
		}
	}
	for num, _ := range hash2 {
		res = append(res, num)
	}
	return res
}
