package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(hasMatch("kvb", "k*v"))
}

func hasMatch(s string, p string) bool {
	splits := strings.Split(p, "*")
	i := 0
	if len(splits[0]) > 0 {
		j := strings.Index(s, splits[0])
		if j == -1 {
			return false
		}
		i = j + len(splits[0])
	}
	s = s[i:]
	if len(splits[1]) > 0 {
		if len(s) == 0 {
			return false
		}
		if strings.Index(s, splits[1]) == -1 {
			return false
		}
	}
	return true
}
