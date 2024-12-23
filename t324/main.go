package main

func main() {
	wiggleSort([]int{1, 5, 1, 1, 6, 4})
}

// https://leetcode.cn/problems/wiggle-sort-ii/
func wiggleSort(nums []int) {

	tmp := make([]int, 5001)
	for _, num := range nums {
		tmp[num]++
	}

	right := 5000
	for i := 1; i < len(nums); i += 2 {
		for tmp[right] == 0 {
			right--
		}
		nums[i] = right
		tmp[right]--
	}

	for i := 0; i < len(nums); i += 2 {
		for tmp[right] == 0 {
			right--
		}
		nums[i] = right
		tmp[right]--
	}
	//fmt.Println(nums)
}
