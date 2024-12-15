package main

func main() {

}

func buttonWithLongestTime(events [][]int) int {
	res, index := events[0][1], events[0][0]
	for i := 1; i < len(events); i++ {
		tmp := events[i][1] - events[i-1][1]
		if tmp == res {
			index = min(index, events[i][0])
		}
		if tmp > res {
			res = tmp
			index = events[i][0]

		}
	}
	return index
}
