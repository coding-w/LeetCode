package main

func main() {

}

// https://leetcode.cn/problems/two-sum-ii-input-array-is-sorted/description/
func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for numbers[l]+numbers[r] != target {
		if numbers[l]+numbers[r] > target {
			r--
		} else {
			l++
		}

	}
	return []int{l + 1, r + 1}
}
