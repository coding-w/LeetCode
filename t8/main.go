package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("+123-45"))
	temp := []int32{50, 69, 74}
	fmt.Println(string(temp))

}

// https://leetcode.cn/problems/string-to-integer-atoi/description/
func myAtoi(s string) int {
	temp := []int32{}
	for _, i := range s {
		if i <= 57 && i >= 48 {
			temp = append(temp, i)
		} else {
			if len(temp) != 0 {
				break
			} else if i == '+' || i == '-' {
				temp = append(temp, i)
			} else if i == ' ' {
				continue
			} else {
				break
			}
		}
	}
	str := string(temp)
	res, _ := strconv.Atoi(str)
	if res < math.MinInt32 {
		res = math.MinInt32
	}
	if res > math.MaxInt32 {
		res = math.MaxInt32
	}
	return res
}
