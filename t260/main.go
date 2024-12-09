package main

import "fmt"

func main() {
	i := 5 & 2
	fmt.Println(i)
	// fmt.Println(2 ^ i)

}

func singleNumber(nums []int) []int {
	// 第一步：对数组中所有数字进行按位异或操作
	xorResult := 0
	for _, num := range nums {
		xorResult ^= num
	}

	// 第二步：找到 xorResult 中最低位为1的位置
	// 因为这个位肯定能区分出两个不同的数字
	rightmostSetBit := xorResult & -xorResult

	// 第三步：根据这个位置将数字分为两组
	var num1, num2 int
	for _, num := range nums {
		if num&rightmostSetBit == 0 {
			num1 ^= num // 属于第一组
		} else {
			num2 ^= num // 属于第二组
		}
	}

	return []int{num1, num2}
}
