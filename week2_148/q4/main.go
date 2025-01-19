package main

import (
	"fmt"
	"sort"
)

const MOD = 1e9 + 7

// 计算曼哈顿距离和
func calculateManhattanSum(m, n, k int) int {
	// 生成所有可能的坐标
	var coordinates []struct{ x, y int }
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			coordinates = append(coordinates, struct{ x, y int }{i, j})
		}
	}

	// 提取横纵坐标
	xCoords := make([]int, len(coordinates))
	yCoords := make([]int, len(coordinates))
	for i, coord := range coordinates {
		xCoords[i] = coord.x
		yCoords[i] = coord.y
	}

	// 对横纵坐标分别排序
	sort.Ints(xCoords)
	sort.Ints(yCoords)

	// 计算横坐标方向的曼哈顿距离
	xSum := 0
	for i := 0; i < len(xCoords); i++ {
		// 横坐标贡献计算
		xSum += (xCoords[i] * i) - (xCoords[i] * (len(xCoords) - 1 - i))
		xSum %= MOD
	}

	// 计算纵坐标方向的曼哈顿距离
	ySum := 0
	for i := 0; i < len(yCoords); i++ {
		// 纵坐标贡献计算
		ySum += (yCoords[i] * i) - (yCoords[i] * (len(yCoords) - 1 - i))
		ySum %= MOD
	}

	// 返回横纵坐标曼哈顿距离之和
	totalSum := (xSum + ySum) % MOD
	return totalSum
}

func main() {
	// 输入
	m := 3
	n := 3
	k := 2

	// 计算所有棋子之间的曼哈顿距离之和
	result := calculateManhattanSum(m, n, k)
	fmt.Println(result) // 输出结果
}
