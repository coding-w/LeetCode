package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	intArray := []int{0, 0}
	fmt.Println(largestNumber(intArray))

}

func largestNumber(nums []int) string {
	strArray := make([]string, len(nums))

	for i, num := range nums {
		strArray[i] = strconv.Itoa(num)
	}
	sort.Slice(strArray, func(i, j int) bool {
		return strArray[i]+strArray[j] > strArray[j]+strArray[i]
	})
	i := 0
	for ; i < len(strArray); i++ {
		if strArray[i] != "0" {
			break
		}
	}
	if i == len(strArray) {
		return "0"
	}
	return strings.Join(strArray, "")
}
