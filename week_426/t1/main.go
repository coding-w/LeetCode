package main

import "fmt"

func main() {
	fmt.Println(smallestNumber(10))
}

func smallestNumber(n int) int {
	c := 0
	for n > 0 {
		n >>= 1
		c++
	}
	return 1<<c - 1

}
