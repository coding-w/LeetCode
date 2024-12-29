package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(answerString("dd", 2))
}

func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}
	m := len(word)
	l := m - numFriends + 1
	var mm uint8 = word[0]
	tmp := make([]int, 0)
	tmp = append(tmp, 0)
	for i := 1; i < m; i++ {
		if mm < word[i] {
			mm = word[i]
			tmp = make([]int, 0)
		}
		if mm == word[i] {
			tmp = append(tmp, i)
		}
	}
	res := make([]string, 0)
	for i := 0; i < len(tmp); i++ {
		if l+tmp[i] > m {
			res = append(res, word[tmp[i]:m])
		} else {
			res = append(res, word[tmp[i]:tmp[i]+l])
		}
	}
	slices.Sort(res)
	return res[len(res)-1]
}
