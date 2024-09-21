package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))

}

// https://leetcode.cn/problems/zigzag-conversion/description/
// tips: 找规律
// todo
func convert(s string, numRows int) string {
	if len(s) < numRows {
		return s
	}
	if numRows == 1 {
		return s
	}
	k := 2 * (numRows - 1)
	res := make([]string, numRows+1)
	for i := 1; i <= numRows; i++ {
		index := i
		var str strings.Builder
		if i == 1 || i == numRows {
			str.WriteByte(s[index-1])
			for index+k-1 < len(s) {
				index += k
				str.WriteByte(s[index-1])
			}
		} else {
			str.WriteByte(s[index-1])
			temp := 2 * (numRows - i)
			for index+temp-1 < len(s) {
				str.WriteByte(s[index+temp-1])
				index += temp
				temp = k - temp
			}
		}
		res[i] = str.String()
	}
	ss := ""
	for i := 1; i <= numRows; i++ {
		ss += res[i]
	}
	return ss
}
