package main

func validSubstringCount(word1 string, word2 string) int64 {
	var count int64 = 0
	if len(word1) < len(word2) {
		return count
	}
	count2 := make([]int, 26)
	for _, v := range word2 {
		count2[v-'a']++
	}
	count1 := make([]int, 26)
	cur := 0
	for i, v := range word1 {
		count1[v-'a']++
		for compare(count1, count2) {
			count += int64(len(word1) - i)
			count1[word1[cur]-'a']--
			cur++
		}
	}
	return count
}

func compare(count1 []int, count2 []int) bool {
	for i := 0; i < 26; i++ {
		if count1[i] < count2[i] {
			return false
		}
	}
	return true
}
