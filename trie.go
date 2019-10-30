package trie

import (
	//"fmt"
)

const asciiSize = 94

// TrieNode is a data structure
type TrieNode struct {
	children   [asciiSize]*TrieNode
	endOfWords bool
}

// GetNode is a function to get a node inside the trie
func GetNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false

	for i := 0; i < asciiSize; i++ {
		node.children[i] = nil
	}

	return node
}

// Insert is a function to insert a key into the trie
func Insert(root *TrieNode, key string) {
	temp := root

	for i := 0; i < len(key); i++ {
		index := key[i] - '!'
		if temp.children[index] == nil {
			temp.children[index] = GetNode()
		}
		temp = temp.children[index]
	}

	temp.endOfWords = true
}

// Search is a function to search for a specific key in the trie
func Search(root *TrieNode, key string) bool {
	temp := root

	for i := 0; i < len(key); i++ {
		index := key[i] - '!'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}
	}

	return (temp != nil && temp.endOfWords)
}

func main() {
	//words := []string{"a", "and", "an", "go", "golang", "man", "mango"}
	//root := GetNode()
	//
	//for i := 0; i < len(words); i++ {
	//	Insert(root, words[i])
	//}
	//
	////fmt.Println("contains words [a]: ", Search(root, "a"))
	////fmt.Println("contains words [mango]: ", Search(root, "mango"))
	////fmt.Println("contains words [lang]: ", Search(root, "lang"))
}
