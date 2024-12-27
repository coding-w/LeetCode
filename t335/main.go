package main

import "fmt"

func main() {
	fmt.Println(isSelfCrossing([]int{3, 2, 2, 3}))
}

// https://leetcode.cn/problems/self-crossing/description/
// 数学归纳法.....做题压根想不到
// https://leetcode.cn/problems/self-crossing/solutions/1072375/tong-ge-lai-shua-ti-la-san-chong-bu-xian-w80r/
// 参考这位大佬的解释，太强了！！！
func isSelfCrossing(distance []int) bool {
	n := len(distance)
	if n < 4 {
		return false
	}
	i := 2
	// 外卷
	for i < n && distance[i] > distance[i-2] {
		i++
	}
	if i == n {
		return false
	}

	// 外卷转内卷， i-1 的长度减去 i-3 的长度
	tmp := 0
	if i >= 4 {
		tmp = distance[i-4]
	}
	if distance[i] >= distance[i-2]-tmp {
		if i >= 3 {
			distance[i-1] -= distance[i-3]
		}
	}

	i++
	// 内卷
	for i < n && distance[i] < distance[i-2] {
		i++
	}

	return i != n

}
