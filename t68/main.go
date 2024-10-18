package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := []string{"This", "is", "an", "example", "of", "text", "justification."}
	str := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do", "123"}
	res := fullJustify(str, 20)
	for i := 0; i < len(res); i++ {
		fmt.Println(res[i])
	}
}

// https://leetcode.cn/problems/text-justification/description/
// tips: 先判断 每一行有哪些单词，再分配格式，最后一行单独处理
func fullJustify(words []string, maxWidth int) []string {
	// 存储结果
	var res []string
	// 存储当前行有哪些单词
	var currentLine []string
	// 存储当前行 单词长度
	count := 0
	for _, word := range words {
		// 检查当前行是否可以添加新单词
		if count+len(word)+len(currentLine) > maxWidth {
			res = append(res, formatLine(currentLine, count, maxWidth, false))
			currentLine = nil
			count = 0
		}
		currentLine = append(currentLine, word)
		count += len(word)
	}
	// 处理最后一行
	if len(currentLine) > 0 {
		res = append(res, formatLine(currentLine, count, maxWidth, true))
	}
	return res
}

// 格式化行的辅助函数
func formatLine(words []string, count, maxWidth int, isLastLine bool) string {
	if isLastLine {
		// 最后一行左对齐
		return leftJustify(words, maxWidth)
	}

	totalSpaces := maxWidth - count
	spaceCount := len(words) - 1
	// 只有一个单词
	if spaceCount == 0 {
		return fmt.Sprintf("%-*s", maxWidth, words[0])
	}

	// 每个单词之间保留的空格数
	spaceSize := totalSpaces / spaceCount
	// 前 extraSpaces 个额外的空格分配
	extraSpaces := totalSpaces % spaceCount

	var sb strings.Builder
	for i, word := range words {
		sb.WriteString(word)
		if i < spaceCount {
			sb.WriteString(strings.Repeat(" ", spaceSize))
			if i < extraSpaces {
				sb.WriteString(" ") // 多余的空格分配
			}
		}
	}
	return sb.String()
}

// 左对齐处理
func leftJustify(words []string, maxWidth int) string {
	var sb strings.Builder
	for _, word := range words {
		sb.WriteString(word + " ")
	}
	res := sb.String()
	res = res[:len(res)-1] // 删除最后一个空格
	return res + strings.Repeat(" ", maxWidth-len(res))
}
