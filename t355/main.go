package main

import (
	"fmt"
	"slices"
)

func main() {
	t := Constructor()
	t.PostTweet(1, 5)
	fmt.Println(t.GetNewsFeed(1))
	t.Follow(1, 2)
	t.PostTweet(2, 6)
	fmt.Println(t.GetNewsFeed(1))
	t.Unfollow(1, 2)
	fmt.Println(t.GetNewsFeed(1))

}

// Twitter
// https://leetcode.cn/problems/design-twitter/
type Twitter struct {
	user    map[int]map[int]bool // 用户关注列表
	article map[int][][]int      // 用户文章列表
	time    int                  // 时间
}

func Constructor() Twitter {
	return Twitter{
		user:    make(map[int]map[int]bool),
		article: make(map[int][][]int),
		time:    0,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.article[userId] = append(this.article[userId], []int{tweetId, this.time})
	this.time++
}

func (this *Twitter) GetNewsFeed(userId int) (res []int) {
	articles := make([][]int, 0)
	if v, ok := this.article[userId]; ok {
		articles = append(articles, v...)
	}
	if v, ok := this.user[userId]; ok {
		for key, _ := range v {
			if v1, ok1 := this.article[key]; ok1 {
				articles = append(articles, v1...)
			}
		}
	}
	slices.SortFunc(articles, func(a, b []int) int {
		return b[1] - a[1]
	})
	for i := 0; i < 10 && i < len(articles); i++ {
		res = append(res, articles[i][0])
	}
	return
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if followeeId != followerId {
		if this.user[followerId] == nil {
			this.user[followerId] = make(map[int]bool)
		}
		this.user[followerId][followeeId] = true
	}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if followeeId != followerId {
		delete(this.user[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
