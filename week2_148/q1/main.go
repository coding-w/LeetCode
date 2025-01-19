package main

func main() {

}

func maxAdjacentDistance(nums []int) int {
	n := len(nums)
	res := abs(nums[n-1] - nums[0])
	for i := 1; i < n; i++ {
		res = max(res, abs(nums[i]-nums[i-1]))
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
