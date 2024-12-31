package main

import "fmt"

func main() {
	nums := []int{3, 5, 4, 6, 2, 1}

	// 1. 冒泡排序
	// fmt.Println(bubbleSort(nums))

	// 2. 选择排序
	// fmt.Println(selectionSort(nums))

	// 3. 插入排序
	// fmt.Println(insertionSort(nums))

	// 4. 希尔排序
	fmt.Println(shellSort(nums))
}

// bubbleSort 冒泡排序
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		for j := 1; j < length-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
	return nums
}

// selectionSort 选择排序
func selectionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		minI := i
		for j := i + 1; j < length; j++ {
			if nums[minI] > nums[j] {
				minI = j
			}
		}
		if minI != i {
			nums[minI], nums[i] = nums[i], nums[minI]
		}
	}
	return nums
}

// insertionSort 插入排序
func insertionSort(nums []int) []int {
	length := len(nums)
	for i := 1; i < length; i++ {
		key := nums[i]
		j := i - 1
		// 将 nums[i] 插入到已排序部分
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = key
	}
	return nums
}

// shellSort 希尔排序
func shellSort(nums []int) []int {
	length := len(nums)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			j := i
			// 将 nums[i] 插入到已排序部分
			for j >= gap && nums[j-gap] > temp {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = temp
		}
	}
	return nums
}
