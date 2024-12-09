package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinimumTime([]int{3, 4, 1}, 1))
}

func findMinimumTime(strength []int, K int) int {

	sort.Ints(strength)
	res := 0
	hmap := make(map[int]int)
	for _, v := range strength {
		hmap[v]++
	}
	x := 1
	tmp := 0
	for i := 0; i < len(strength); i++ {
		tmp = x
		v, ok := hmap[tmp]
		if ok && v == 0 {
			continue
		}
		res++
		for tmp-strength[i] < 0 {
			res++
			tmp += x
			_, ok = hmap[tmp]
			if ok {
				x += K
				tmp = 0
			}
		}
		x += K
		tmp = 0
	}
	return res

}
