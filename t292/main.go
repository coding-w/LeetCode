package main

import "fmt"

func main() {
	fmt.Println(canWinNim(12))
}

// https://leetcode.cn/problems/nim-game/
func canWinNim(n int) bool {
	return n%4 != 0
}
