// -*- compile-command: "go run main.go"; -*-

package main

import (
	"sort"
)

type WordSequence struct {
	StartOffset int
	EndOffset   int
	WordCount   int
	Words       []string
}

func findWordSequences(text string, words map[string]struct{}) []*WordSequence {
	if len(text) == 0 || len(words) == 0 {
		return nil
	}

	// Build trie for efficient prefix matching
	trie := buildTrie(words)

	var sequences []*WordSequence
	textLen := len(text)

	// For each starting position, find the longest sequence of words
	for i := 0; i < textLen; i++ {
		sequence := findSequenceAtPosition(text, i, trie)
		if sequence != nil && sequence.WordCount >= 2 {
			sequences = append(sequences, sequence)
		}
	}

	// Sort by word count descending, then by start position ascending for stability
	sort.Slice(sequences, func(i, j int) bool {
		if sequences[i].WordCount != sequences[j].WordCount {
			return sequences[i].WordCount > sequences[j].WordCount
		}
		return sequences[i].StartOffset < sequences[j].StartOffset
	})

	return sequences
}

type trieNode struct {
	children map[byte]*trieNode
	isWord   bool
	word     string
}

func buildTrie(words map[string]struct{}) *trieNode {
	root := &trieNode{children: make(map[byte]*trieNode)}

	for word := range words {
		node := root
		for i := 0; i < len(word); i++ {
			c := word[i]
			if node.children[c] == nil {
				node.children[c] = &trieNode{children: make(map[byte]*trieNode)}
			}
			node = node.children[c]
		}
		node.isWord = true
		node.word = word
	}

	return root
}

func findSequenceAtPosition(text string, startPos int, trie *trieNode) *WordSequence {
	var foundWords []string
	pos := startPos
	textLen := len(text)

	for pos < textLen {
		word, endPos := findWordAtPosition(text, pos, trie)
		if word == "" {
			break
		}
		foundWords = append(foundWords, word)
		pos = endPos
	}

	if len(foundWords) < 2 {
		return nil
	}

	return &WordSequence{
		StartOffset: startPos,
		EndOffset:   pos,
		WordCount:   len(foundWords),
		Words:       foundWords,
	}
}

func findWordAtPosition(text string, pos int, trie *trieNode) (string, int) {
	node := trie
	textLen := len(text)
	lastWordEnd := -1
	lastWord := ""

	for i := pos; i < textLen; i++ {
		c := text[i]
		if node.children[c] == nil {
			break
		}
		node = node.children[c]
		if node.isWord {
			lastWord = node.word
			lastWordEnd = i + 1
		}
	}

	if lastWordEnd == -1 {
		return "", pos
	}

	return lastWord, lastWordEnd
}
