package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(11))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	var i, j int
	y := strconv.Itoa(x)
	length := len(y)
	if length%2 == 0 {
		i = length/2 - 1
		j = length / 2
	} else {
		i = length/2 - 1
		j = length/2 + 1
	}
	for i >= 0 && j < length {
		if y[i] != y[j] {
			return false
		}
		i--
		j++
	}
	return true
}
