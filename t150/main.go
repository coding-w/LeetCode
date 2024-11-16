package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

// https://leetcode.cn/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	//opr := make([]string, 0)
	res := make([]int, 0)
	m := make(map[string]int)
	m["*"] = 2
	m["/"] = 2
	m["+"] = 1
	m["-"] = 1
	for _, token := range tokens {
		num, err := strconv.Atoi(token)
		if err == nil {
			res = append(res, num)
		} else {
			res = compute(token, res)
			//if len(opr) == 0 {
			//	res = compute(token, res)
			//} else {
			//	if m[token] > m[opr[0]] {
			//		res = compute(token, res)
			//	}
			//
			//}
		}
	}
	return res[0]

}

func compute(token string, res []int) []int {
	n := len(res)
	switch token {
	case "+":
		res[n-2] += res[n-1]
		break
	case "-":
		res[n-2] -= res[n-1]
		break
	case "*":
		res[n-2] *= res[n-1]
		break
	case "/":
		res[n-2] /= res[n-1]
		break
	}
	return res[:n-1]
}
