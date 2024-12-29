package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(rankTeams([]string{"FVSHJIEMNGYPTQOURLWCZKAX", "AITFQORCEHPVJMXGKSLNZWUY", "OTERVXFZUMHNIYSCQAWGPKJL", "VMSERIJYLZNWCPQTOKFUHAXG", "VNHOZWKQCEFYPSGLAMXJIUTR", "ANPHQIJMXCWOSKTYGULFVERZ", "RFYUXJEWCKQOMGATHZVILNSP", "SCPYUMQJTVEXKRNLIOWGHAFZ", "VIKTSJCEYQGLOMPZWAHFXURN", "SVJICLXKHQZTFWNPYRGMEUAO", "JRCTHYKIGSXPOZLUQAVNEWFM", "NGMSWJITREHFZVQCUKXYAPOL", "WUXJOQKGNSYLHEZAFIPMRCVT", "PKYQIOLXFCRGHZNAMJVUTWES", "FERSGNMJVZXWAYLIKCPUQHTO", "HPLRIUQMTSGYJVAXWNOCZEKF", "JUVWPTEGCOFYSKXNRMHQALIZ", "MWPIAZCNSLEYRTHFKQXUOVGJ", "EZXLUNFVCMORSIWKTYHJAQPG", "HRQNLTKJFIEGMCSXAZPYOVUW", "LOHXVYGWRIJMCPSQENUAKTZF", "XKUTWPRGHOAQFLVYMJSNEIZC", "WTCRQMVKPHOSLGAXZUEFYNJI"}))
	//fmt.Println(rankTeams([]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}))
}

// https://leetcode.cn/problems/rank-teams-by-votes/?envType=daily-question&envId=2024-12-29
// 每日一题

func rankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}

	// 初始化二维数组：每个队伍统计每个位置的得票情况，最后一位存字母编号
	tmp := make([][]int, 26)
	for i := range tmp {
		tmp[i] = make([]int, len(votes[0])+2)
		tmp[i][len(votes[0])] = i
	}

	for _, vote := range votes {
		for i := 0; i < len(vote); i++ {
			tmp[vote[i]-'A'][i]++
			// 统计每个队伍在每个位置的得票数
			tmp[vote[i]-'A'][len(votes[0])+1] = 1
		}
	}

	// 自定义排序规则
	slices.SortFunc(tmp, func(a, b []int) int {
		// 按每个位置的票数从高到低排序
		for i := 0; i < len(votes[0]); i++ {
			if a[i] != b[i] {
				return b[i] - a[i] // 按票数降序排序
			}
		}
		// 如果票数完全相同，按队伍字母顺序升序
		return a[len(votes[0])] - b[len(votes[0])]
	})

	// 构造结果字符串
	var res strings.Builder
	for i := 0; i < 26; i++ {
		if tmp[i][len(votes[0])+1] != 0 { // 只有被投票的队伍才加入结果
			res.WriteRune(rune('A' + tmp[i][len(votes[0])]))
		}
	}

	return res.String()
}
