package main

import "fmt"

func main() {
	fmt.Println(generateKey(12, 123, 1234))
}

func generateKey(num1 int, num2 int, num3 int) int {
	res := 0
	res += min(num1/1000, num2/1000, num3/1000) * 1000
	num1 %= 1000
	num2 %= 1000
	num3 %= 1000
	res += min(num1/100, num2/100, num3/100) * 100
	num1 %= 100
	num2 %= 100
	num3 %= 100
	res += min(num1/10, num2/10, num3/10) * 10
	num1 %= 10
	num2 %= 10
	num3 %= 10
	res += min(num1, num2, num3)
	return res
}
