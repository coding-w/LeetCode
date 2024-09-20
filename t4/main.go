package main

import "fmt"

func main() {

	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}

// https://leetcode.cn/problems/median-of-two-sorted-arrays/submissions/540457953/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	res := make([]int, n+m)
	i := 0
	j := 0
	for i < n && j < m {
		if nums1[i] <= nums2[j] {
			res[i+j] = nums1[i]
			i++
		} else {
			res[i+j] = nums2[j]
			j++
		}
	}
	if i < n {
		for i < n {
			res[i+j] = nums1[i]
			i++
		}
	}
	if j < m {
		for j < m {
			res[i+j] = nums2[j]
			j++
		}
	}
	if len(res)%2 == 0 {
		return float64(res[len(res)/2]+res[len(res)/2-1]) / 2
	} else {
		return float64(res[len(res)/2])
	}
}
