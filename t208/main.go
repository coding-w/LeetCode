package main

func main() {
	constructor := Constructor()
	constructor.Insert("apple")
	constructor.Insert("appasdasd")
	println(constructor.Search("apple"))
	println(constructor.StartsWith("app"))

}

type TrieNode struct {
	child map[rune]*TrieNode
	isEnd bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{root: &TrieNode{child: make(map[rune]*TrieNode)}}
}

func (this *Trie) Insert(word string) {
	node := this.root
	for _, char := range word {
		if _, exists := node.child[char]; !exists {
			node.child[char] = &TrieNode{child: make(map[rune]*TrieNode)}
		}
		node = node.child[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.searchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

func (t *Trie) searchPrefix(prefix string) *TrieNode {
	node := t.root
	for _, char := range prefix {
		if _, exists := node.child[char]; !exists {
			return nil
		}
		node = node.child[char]
	}
	return node
}
