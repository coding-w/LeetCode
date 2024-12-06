package main

import "fmt"

func main() {
	fmt.Println(addDigits(10))
}

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	res := 0
	for num >= 10 {
		res += num % 10
		num = num / 10
	}
	res += num
	return addDigits(res)
}
