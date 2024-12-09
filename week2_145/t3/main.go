package main

import (
	"fmt"
	"math"
	"strconv"
)

// 判断一个数是否为质数
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 计算从n变到m的最小代价
func minCost(n, m int) int {
	// 将n和m转换为字符串，方便逐位处理
	//nStr := strconv.Itoa(n)
	//mStr := strconv.Itoa(m)

	// 如果n和m相同，直接返回0
	if n == m {
		return 0
	}

	// BFS的队列存储状态，队列中的元素是当前数字和当前代价
	queue := [][]int{{n, 0}}
	visited := map[int]bool{n: true} // 记录访问过的数字，避免重复

	// BFS搜索最小代价
	for len(queue) > 0 {
		// 获取当前状态
		current := queue[0][0]
		cost := queue[0][1]
		queue = queue[1:] // 出队

		// 如果当前数字是质数，跳过
		if isPrime(current) {
			continue
		}

		// 将当前数字转换为字符串，逐位操作
		currentStr := strconv.Itoa(current)

		// 枚举每一位的增减操作
		for i := 0; i < len(currentStr); i++ {
			// 枚举增减1的操作
			for _, delta := range []int{-1, 1} {
				newDigit := int(currentStr[i]-'0') + delta

				// 如果数字超出范围[0, 9]，则跳过
				if newDigit < 0 || newDigit > 9 {
					continue
				}

				// 修改当前数字的第i位
				newNumber := currentStr[:i] + strconv.Itoa(newDigit) + currentStr[i+1:]
				newNum, _ := strconv.Atoi(newNumber)

				// 如果新状态未访问过，且不是质数，则入队
				if !visited[newNum] && !isPrime(newNum) {
					visited[newNum] = true
					// 计算新状态的代价
					newCost := cost + newNum
					// 如果达到了目标m，返回最小代价
					if newNum == m {
						return newCost
					}
					// 将新状态入队
					queue = append(queue, []int{newNum, newCost})
				}
			}
		}
	}

	// 如果无法将n转化为m，返回-1
	return -1
}

func main() {
	n := 12
	m := 21
	fmt.Println(minCost(n, m)) // 输出最小代价
}
