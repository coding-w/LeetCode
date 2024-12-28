package main

func main() {

}

// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
func intersect(nums1 []int, nums2 []int) []int {
	h1 := make(map[int][]int)
	h2 := make(map[int][]int)
	for _, num := range nums1 {
		h1[num] = append(h1[num], num)

	}
	for _, num := range nums2 {
		h2[num] = append(h2[num], num)
	}
	res := make([]int, 0)
	for k, v := range h1 {
		v2 := h2[k]
		if len(v2) > 0 {
			if len(v2) == len(v) {
				res = append(res, v...)
			} else if len(v) > len(v2) {
				res = append(res, v2...)
			} else {
				res = append(res, v...)
			}
		}
	}
	return res
}
