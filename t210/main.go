package main

import "fmt"

func main() {
	prerequisites := [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 3}}
	fmt.Println(findOrder(5, prerequisites))
}

// https://leetcode.cn/problems/course-schedule-ii/
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, item := range prerequisites {
		curr, pre := item[0], item[1]
		graph[pre] = append(graph[pre], curr)
		indegree[curr]++
	}

	var queue []int
	var res []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		for _, neighbor := range graph[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}
	if len(res) != numCourses {
		return []int{}
	}
	return res
}
