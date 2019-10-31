# Text Editor Re-implementation in Go

Created by Hanyu Li  

----
## Introduction
This file serves as a README for the re-implemented Text Editor to describe the following:   
1. Any design decisions or tradeoffs made by the author  
2. Any extensions added or the author intended to add  
3. How to run the program

----
## Design Decisions  
### Data Structure for Document

Using a simple string to store the document is intuitive but low in efficiency. In Go, string is a read-only slice of byte, and manipulating or concatenating strings could consume a lot of time. In this re-implementation, a data structure called [Rope](https://en.wikipedia.org/wiki/Rope_%28data_structure%29) is used to efficiently store and manipulate large text.  

>A rope is a binary tree where each leaf (end node) holds a string and a length (also known as a "weight"), and each node further up the tree holds the sum of the lengths of all the leaves in its left subtree. A node with two children thus divides the whole string into two parts: the left subtree stores the first part of the string, the right subtree stores the second part of the string, and node's weight is the sum of the left child's weight along with all of the nodes contained in its subtree. 

Since a rope is actually a binary tree, now the time complexity for searching is reduced to O(logn), and manipulating becomes faster because there will be no more type conversion in cut or paste operation.  This drastically decreases the average running time for operation **CutPaste** and **CopyPaste**.  
The price of using a rope to store the document is that when we want to get the text as **string**, type conversion is required, and thus increases the running time for function **GetText()**. To give this function a better performance, another string is used to store the text, only being called in **GetText()**, and everytime the document is changed(either cut or paste operation is done), this string will get updated. The tradeoff is increasing the running time of **Cut()** and **Paste()** since now they have to perform type conversion.

### Data Structure for Dictionary  

The theoretical time complexity of search in a map is O(1). However, as a result of dealing with hash collision when the number of keys increases, there will be a significant drop in the speed of search. This makes map a less ideal data structure to store a huge number of strings.  
An idea similar to document storage described above could be applied here. Instead of sticking with map, we use [Trie](https://en.wikipedia.org/wiki/Trie) in  Go's **collections** package for the dictionary.  

>In computer science, a trie, also called digital tree or prefix tree, is a kind of search tree—an ordered tree data structure used to store a dynamic set or associative array where the keys are usually strings. Unlike a binary search tree, no node in the tree stores the key associated with that node; instead, its position in the tree defines the key with which it is associated. All the descendants of a node have a common prefix of the string associated with that node, and the root is associated with the empty string. Keys tend to be associated with leaves, though some inner nodes may correspond to keys of interest. Hence, keys are not necessarily associated with every node.  

The trie implementation used for this project is from [this website](http://www.code2succeed.com/golang-insert-and-search-trie/).

----
## Extensions  
### Additional Function

Another benefit of using a rope to store the document is that insert operation would be quite cheap, so this new function **Insert(i int, str string)** is added to the re-implementation.

### Further Steps

Materials from the Internet suggest that a widely used data structure in real text editors is [Piece Table](https://en.wikipedia.org/wiki/Piece_table). Although implementing a piece table from scratch could be time consuming, it is still a good direction to explore.

----
## Tips for Running the Program
1. Run `go get github.com/vinzmay/go-rope` to install package rope
2. Make sure you have the file **trie.go** as custom package to import

----
## Test Results Comparison
See the results in the txt file [here](https://github.com/lizzielee/Neeva-Text-Editor-Project/blob/master/TestResults.txt)
