package editor_test

import (
	"bufio"
	"github.com/vinzmay/go-rope"
	"os"
	"strconv"
	"strings"
	"testing"
	"trie"
)

// TextEditor provides functionality for manipulating and analyzing text documents.
type TextEditor interface {
	// Removes characters [i..j) from the document and places them in the clipboard.
	// Previous clipboard contents is overwritten.
	Cut(i, j int)
	// Places characters [i..j) from the document in the clipboard.
	// Previous clipboard contents is overwritten.
	Copy(i, j int)
	// Inserts the contents of the clipboard into the document starting at position i.
	// Nothing is inserted if the clipboard is empty.
	Paste(i int)
	// Returns the document as a string.
	GetText() string
	// Returns the number of misspelled words in the document. A word is considered misspelled
	// if it does not appear in /usr/share/dict/words or any other dictionary (of comparable size)
	// that you choose.
	Misspellings() int
}

type SimpleEditor struct {
	staticDocument string
	document   rope.Rope
	dictionary trie.TrieNode
	pasteText  string
	builder strings.Builder
}

func NewSimpleEditor(document string) TextEditor {
	fileHandle, _ := os.Open("/usr/share/dict/words")
	defer fileHandle.Close()
	dict := trie.GetNode()
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		currentStr := scanner.Text()
		trie.Insert(dict, currentStr)
	}
	var builder strings.Builder
	return &SimpleEditor{staticDocument: document, document: *rope.New(document),
		dictionary: *dict, builder: builder}
}

func (s *SimpleEditor) Cut(i, j int) {
	//s.pasteText = s.document.Substr(i,j).String()
	leftPart, tmp := s.document.Split(i)
	pasteText, rightPart := tmp.Split(j)
	s.pasteText = pasteText.String()

	s.document = *leftPart.Concat(rightPart)
	s.staticDocument = s.document.String()
}

func (s *SimpleEditor) Copy(i, j int) {
	s.pasteText = s.document.Substr(i, j).String()
}

func (s *SimpleEditor) Paste(i int) {

	s.document = *s.document.Insert(i, s.pasteText)
}

func (s *SimpleEditor) GetText() string {
	return s.staticDocument
}

func (s *SimpleEditor) Misspellings() int {
	result := 0
	visited := make(map[string]bool)

	for _, word := range strings.Fields(s.staticDocument) {
		if !visited[word] {
			visited[word] = true
			if !trie.Search(&s.dictionary, word) {
				result++
			}
		}
	}
	return result
}

func (s *SimpleEditor) Insert(i int, str string) {
	s.document = *s.document.Insert(i, str)
}

func BenchmarkClipboard(b *testing.B) {
	cases := []struct {
		data string
	}{
		{strings.Repeat("Neeva is awesome!", 10)},
		{strings.Repeat("Neeva is awesome!", 100)},
	}
	for _, tc := range cases {
		s := NewSimpleEditor(tc.data)
		b.Run("CutPaste"+strconv.Itoa(len(tc.data)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				if n%2 == 0 {
					s.Cut(1, 3)
				} else {
					s.Paste(2)
				}
			}
		})
		s = NewSimpleEditor(tc.data)
		b.Run("CopyPaste"+strconv.Itoa(len(tc.data)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				if n%2 == 0 {
					s.Copy(1, 3)
				} else {
					s.Paste(2)
				}
			}
		})
		s = NewSimpleEditor(tc.data)
		b.Run("GetText"+strconv.Itoa(len(tc.data)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = s.GetText()
			}
		})
		s = NewSimpleEditor(tc.data)
		b.Run("Misspellings"+strconv.Itoa(len(tc.data)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_ = s.Misspellings()
			}
		})
	}
}
