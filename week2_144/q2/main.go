package main

import "fmt"

func main() {
	fmt.Println(shiftDistance("leet", "code",
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}))
}

func shiftDistance(s string, t string, nextCost []int, previousCost []int) int64 {
	n := len(s)
	var total int64 = 0

	for i := 0; i < n; i++ {
		cur := int(s[i] - 'a')
		target := int(t[i] - 'a')

		r := 0
		steps := (target - cur + 26) % 26
		for step := 0; step < steps; step++ {
			r += nextCost[(cur+step)%26]
		}

		l := 0
		steps = (cur - target + 26) % 26
		for step := 0; step < steps; step++ {
			l += previousCost[(cur-step+26)%26]
		}

		total += int64(min(r, l))
	}

	return total
}
