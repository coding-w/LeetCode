package main

import "fmt"

func main() {
	constructor := Constructor()
	//constructor.AddWord("bad")
	//constructor.AddWord("dad")
	//constructor.AddWord("mad")
	constructor.AddWord("a")

	//fmt.Println(constructor.Search("pad"))
	//fmt.Println(constructor.Search("bad"))
	fmt.Println(constructor.Search(".a"))
	//fmt.Println(constructor.Search(".ad."))
}

type TrieNode struct {
	child map[rune]*TrieNode
	isEnd bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{child: make(map[rune]*TrieNode)}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for _, char := range word {
		if _, exists := node.child[char]; !exists {
			node.child[char] = &TrieNode{child: make(map[rune]*TrieNode)}
		}
		node = node.child[char]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this.root
	for index, char := range word {
		if char == '.' {
			return this.Find(word[index:], node)
		} else {
			if _, exists := node.child[char]; !exists {
				return false
			}
			node = node.child[char]
		}
	}
	return node.isEnd
}

func (this *WordDictionary) Find(word string, node *TrieNode) bool {
	if len(word) == 0 {
		return node.isEnd
	}
	char := word[0]
	if char == '.' {
		if len(node.child) == 0 {
			return false
		}
		for _, child := range node.child {
			if this.Find(word[1:], child) {
				return true
			}
		}
	} else {
		if v, exists := node.child[rune(char)]; exists {
			return this.Find(word[1:], v)
		}
	}
	return false
}
