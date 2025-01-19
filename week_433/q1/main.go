package main

func main() {

}

func subarraySum(nums []int) int {

	res := 0
	for i := 0; i < len(nums); i++ {
		start := max(0, i-nums[i])
		res += add(nums[start : i+1])
	}
	return res

}

func add(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}
