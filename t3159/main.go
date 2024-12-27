package main

func main() {

}

// https://leetcode.cn/problems/find-occurrences-of-an-element-in-an-array/
// 每日一题
func occurrencesOfElement(nums []int, queries []int, x int) (ans []int) {
	tmp := make([]int, 0)
	for i, v := range nums {
		if v == x {
			tmp = append(tmp, i)
		}
	}
	n := len(tmp)
	for _, v := range queries {
		if v <= n {
			ans = append(ans, tmp[v-1])
		} else {
			ans = append(ans, -1)
		}

	}
	return
}
