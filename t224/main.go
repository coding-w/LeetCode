package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//fmt.Println(calculate("(1+(-4+5+2)-3)+(6+18)"))
	//fmt.Println(calculate("1-(     -2)"))
	fmt.Println(calculate("- (3 + (4 + 5))"))
	fmt.Println(calculate("-(-2)+4"))
	//fmt.Println(calculate("1 + 1"))
}

func calculate(s string) int {
	s = strings.ReplaceAll(s, " ", "")
	if s[0] == '-' {
		s = "0" + s
	}
	res, _ := recursion(s, 0)
	return res
}

func recursion(s string, index int) (int, int) {
	res := 0
	var tmp int
	if s[index] == '-' {
		res, index = recursion(s, index+1)
		return -res, index
	}
	if s[index] == '(' {
		res, index = recursion(s, index+1)
		if index == len(s) {
			return res, index
		}

	}
	if s[index] >= '0' && s[index] <= '9' {
		tmp, index = getNum(s, index)
		res += tmp
	}
	for index < len(s) {
		if s[index] == '(' {
			index++
			tmp, index = recursion(s, index)
			res += tmp
			continue
		}
		if s[index] == ')' {
			return res, index + 1
		}
		switch s[index] {
		case '-':
			if s[index+1] == '(' {
				tmp, index = recursion(s, index+2)
				res -= tmp
				continue
			} else {
				tmp, index = getNum(s, index+1)
				res -= tmp
			}
			break
		case '+':
			if s[index+1] == '(' {
				tmp, index = recursion(s, index+2)
				res += tmp
				continue
			} else {
				tmp, index = getNum(s, index+1)
				res += tmp
			}
			break
		default:
			index++
		}
	}
	return res, index
}

func getNum(s string, index int) (int, int) {
	start := index
	for index < len(s) && s[index] >= '0' && s[index] <= '9' {
		index++
	}
	num, _ := strconv.Atoi(s[start:index])
	return num, index
}
