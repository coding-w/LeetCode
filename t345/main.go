package main

func main() {

}

// https://leetcode.cn/problems/reverse-vowels-of-a-string/
// 题目真难读懂！！！
func reverseVowels(s string) string {
	hash := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}

	str := []byte(s)

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		for i < len(str) && !hash[str[i]] {
			i++
		}
		for j >= 0 && !hash[str[j]] {
			j--
		}
		if i < j {
			str[i], str[j] = str[j], str[i]
		}
	}
	return string(str)
}
