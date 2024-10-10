package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

// https://leetcode.cn/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string][]string)
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedStr := string(runes)
		resMap[sortedStr] = append(resMap[sortedStr], str)
	}
	res := make([][]string, 0)
	for _, str := range resMap {
		res = append(res, str)
	}
	return res
}
