package main

import (
	"fmt"
	"math/rand"
	"time"
)

// main.go golang 实现
func main() {
	// 1. 冒泡排序
	fmt.Println(bubbleSort([]int{3, 5, 4, 6, 2, 1}))

	// 2. 选择排序
	// fmt.Println(selectionSort([]int{3, 5, 4, 6, 2, 1}))

	// 3. 插入排序
	// fmt.Println(insertionSort([]int{3, 5, 4, 6, 2, 1}))

	// 4. 希尔排序
	// fmt.Println(shellSort([]int{3, 5, 4, 6, 2, 1}))

	// 5. 归并排序
	// fmt.Println(mergeSort([]int{3, 5, 4, 6, 2, 1}))

	// 6. 快速排序
	// fmt.Println(quickSort([]int{3, 5, 4, 6, 2, 1}))

	// 7. 堆排序
	//fmt.Println(heapSort([]int{3, 5, 4, 6, 2, 1}))

	// 8. 计数排序
	// fmt.Println(countingSort([]int{3, 5, 4, 6, 2, 1}))

	// 9. 桶排序
	// fmt.Println(bucketSort([]int{3, 5, 4, 6, 2, 1}))

	// 10. 基数排序
	//fmt.Println(radixSort([]int{113, 112, 210, 6}))

	// 11. Timsort
	//fmt.Println(timSort([]int{3, 5, 4, 6, 2, 1}))

	// 12. 猴子排序
	fmt.Println(bogoSort([]int{3, 5, 4, 6, 2, 1}))
}

// bubbleSort 冒泡排序
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		swapped := false
		for j := 1; j < length-i; j++ {
			if nums[j-1] > nums[j] {
				nums[j-1], nums[j] = nums[j], nums[j-1]
				swapped = true
			}
		}
		if !swapped {
			break
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
		temp := nums[i]
		j := i - 1
		// 将 nums[i] 插入到已排序部分
		for j >= 0 && nums[j] > temp {
			nums[j+1] = nums[j]
			j = j - 1
		}
		nums[j+1] = temp
	}
	return nums
}

// shellSort 希尔排序
func shellSort(nums []int) []int {
	length := len(nums)
	for gap := length / 2; gap > 0; gap /= 2 {
		for i := gap; i < length; i++ {
			temp := nums[i]
			j := i - gap
			// 将 nums[i] 插入到已排序部分
			for j >= 0 && nums[j] > temp {
				nums[j+gap] = nums[j]
				j = j - gap
			}
			nums[j+gap] = temp
		}
	}
	return nums
}

// mergeSort 归并排序
func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

// merge 合并两个有序数组
func merge(left, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	// 将剩余的元素添加到结果中
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)
	return merged
}

// quickSort 快速排序
func quickSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	pivot := nums[0]
	var less []int
	var greater []int
	for i := 1; i < length; i++ {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}
	return append(append(quickSort(less), pivot), quickSort(greater)...)
}

// heapSort 堆排序
func heapSort(nums []int) []int {
	length := len(nums)

	// 构建最大堆
	for i := length/2 - 1; i >= 0; i-- {
		heapify(nums, length, i)
	}

	// 一个个从堆中取出元素
	for i := length - 1; i > 0; i-- {
		// 交换当前根节点与末尾元素
		nums[0], nums[i] = nums[i], nums[0]
		// 调整堆
		heapify(nums, i, 0)
	}

	return nums
}

// heapify 调整堆
func heapify(nums []int, heapSize, rootIndex int) {
	largest := rootIndex
	leftChild := 2*rootIndex + 1
	rightChild := 2*rootIndex + 2

	// 如果左子节点大于根节点
	if leftChild < heapSize && nums[leftChild] > nums[largest] {
		largest = leftChild
	}

	// 如果右子节点大于最大节点
	if rightChild < heapSize && nums[rightChild] > nums[largest] {
		largest = rightChild
	}

	// 如果最大节点不是根节点
	if largest != rootIndex {
		nums[rootIndex], nums[largest] = nums[largest], nums[rootIndex]
		// 递归地调整受影响的子树
		heapify(nums, heapSize, largest)
	}
}

// countingSort 计数排序
func countingSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 找到数组中的最大值和最小值
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	// 创建计数数组
	count := make([]int, maxNum-minNum+1)

	// 统计每个元素的出现次数
	for _, num := range nums {
		count[num-minNum]++
	}

	// 生成排序后的数组
	index := 0
	for i, c := range count {
		for c > 0 {
			nums[index] = i + minNum
			index++
			c--
		}
	}

	return nums
}

// bucketSort 桶排序
func bucketSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 找到数组中的最大值和最小值
	maxNum, minNum := nums[0], nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
		if num < minNum {
			minNum = num
		}
	}

	// 创建桶
	bucketCount := len(nums)
	buckets := make([][]int, bucketCount)

	// 将元素分配到桶中
	for _, num := range nums {
		index := (num - minNum) * (bucketCount - 1) / (maxNum - minNum)
		buckets[index] = append(buckets[index], num)
	}

	// 对每个桶进行排序并合并结果
	sortedNums := make([]int, 0, len(nums))
	for _, bucket := range buckets {
		sortedNums = append(sortedNums, mergeSort(bucket)...)
	}

	return sortedNums
}

// radixSort 基数排序
func radixSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	// 找到数组中的最大值
	maxNum := nums[0]
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	// 从个位开始，对每一位进行计数排序
	exp := 1
	for maxNum/exp > 0 {
		nums = countingSortByDigit(nums, exp)
		exp *= 10
	}

	return nums
}

// countingSortByDigit 对数组的某一位进行计数排序
func countingSortByDigit(nums []int, exp int) []int {
	output := make([]int, len(nums))
	count := make([]int, 10)

	// 统计每个数字出现的次数
	for _, num := range nums {
		index := (num / exp) % 10
		count[index]++
	}

	// 计算每个数字的累积计数
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	// 根据当前位的数字，将元素放入输出数组
	for i := len(nums) - 1; i >= 0; i-- {
		index := (nums[i] / exp) % 10
		output[count[index]-1] = nums[i]
		count[index]--
	}

	return output
}

// timSort
func timSort(nums []int) []int {
	const run = 32
	n := len(nums)

	// 对每个小块使用插入排序
	for i := 0; i < n; i += run {
		insertionSortRange(nums, i, min(i+run, n))
	}

	// 合并块
	for size := run; size < n; size *= 2 {
		for left := 0; left < n; left += 2 * size {
			mid := min(left+size, n)
			right := min(left+2*size, n)
			if mid < right {
				mergeRange(nums, left, mid, right)
			}
		}
	}

	return nums
}

// insertionSortRange 对指定范围使用插入排序
func insertionSortRange(nums []int, left, right int) {
	for i := left + 1; i < right; i++ {
		temp := nums[i]
		j := i - 1
		for j >= left && nums[j] > temp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = temp
	}
}

// mergeRange 合并两个有序的子数组
func mergeRange(nums []int, left, mid, right int) {
	leftPart := append([]int(nil), nums[left:mid]...)
	rightPart := append([]int(nil), nums[mid:right]...)

	i, j, k := 0, 0, left
	for i < len(leftPart) && j < len(rightPart) {
		if leftPart[i] <= rightPart[j] {
			nums[k] = leftPart[i]
			i++
		} else {
			nums[k] = rightPart[j]
			j++
		}
		k++
	}

	for i < len(leftPart) {
		nums[k] = leftPart[i]
		i++
		k++
	}

	for j < len(rightPart) {
		nums[k] = rightPart[j]
		j++
		k++
	}
}

// bogoSort 猴子排序
func bogoSort(nums []int) []int {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for !isSorted(nums) {
		shuffle(nums)
	}

	return nums
}

// isSorted 检查数组是否有序
func isSorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

// shuffle 随机打乱数组
func shuffle(nums []int) {
	for i := range nums {
		j := rand.Intn(len(nums))
		nums[i], nums[j] = nums[j], nums[i]
	}
}
