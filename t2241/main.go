package main

import (
	"fmt"
	"sort"
)

func main() {
	c := Constructor()
	c.Deposit([]int{0, 0, 1, 2, 1})
	fmt.Println(c.Withdraw(600))
	c.Deposit([]int{0, 1, 0, 1, 1})
	sort.Ints([]int{})
	fmt.Println(c.Withdraw(600))
	fmt.Println(c.Withdraw(550))
}

// ATM
// https://leetcode.cn/problems/design-an-atm-machine/?envType=daily-question&envId=2025-01-05
type ATM struct {
	count [5]int
	tmp   [5]int
}

func Constructor() ATM {
	return ATM{
		count: [5]int{},
		tmp:   [5]int{20, 50, 100, 200, 500},
	}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < 5; i++ {
		this.count[i] += banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		num := amount / this.tmp[i]
		if num > 0 {
			num = min(num, this.count[i])
			amount -= num * this.tmp[i]
		}
		res[i] = num
		if amount == 0 {
			for j := 0; j < 5; j++ {
				this.count[j] -= res[j]
			}
			return res
		}
	}
	return []int{-1}
}
