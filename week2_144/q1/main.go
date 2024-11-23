package main

import "fmt"

func main() {
	fmt.Println(canAliceWin(10))

}

func canAliceWin(n int) bool {
	if n < 10 {
		return false
	}
	a := 10
	for n > 0 {

		if n-a < 0 {
			return false
		} else {
			n -= a
			a--
		}
		if n-a < 0 {
			return true
		} else {
			n -= a
			a--
		}
	}
	return false
}
