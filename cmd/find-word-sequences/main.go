// -*- compile-command: "go run main.go ../../skips-77777/skip-77777/skip-77777-00*.txt"; -*-

// find-word-sequences searches a string of ASCII lower-case letters
// for consecutive words and returns the results.
// claude.ai was asked to write the initial algorithm, and then I have modified
// it to generate the `trie` once before calling `findWordSequences` multiple times
// with different text (generated with different offsets).
//
// Original prompt:
//
//	Using the Go programming language, write a single function that uses the most
//	efficient algorithm for the following requirements:
//	* the function signature is: `func findWordSequences(text string, words map[string]struct{}) []*WordSequence`
//	* the `text` param strictly consists of ASCII lower-case letters from 'a' to 'z' and nothing
//	  else - no spaces, no punctuation, nothing else - and its length can be up to 3 million characters
//	* the `words` map consists of 12828 lower-case words
//	* the function returns all word sequences of 2-or-more words, and the `WordSequence` structs are
//	  sorted (descending) by the total number of immediately-sequential words found
//	* the `WordSequence` struct must contain the starting and ending offsets of the sequence of words,
//	  the number of words found, and the actual sequence of words found (as a slice of strings)
//
// claude.ai wrote the following:
// I'll write an efficient Go function to find word sequences in text.
// This requires careful algorithm design given the large input sizes (3M characters, 12K+ words).
// This implementation uses several optimization techniques for efficiency:
//
// **Algorithm Design:**
//  1. **Trie Data Structure**: Built from the word dictionary for O(m) prefix matching where m is
//     the word length, much faster than checking every word
//  2. **Greedy Longest Match**: At each position, finds the longest possible word to maximize sequence length
//  3. **Single Pass**: Scans each position once, avoiding redundant work
//
// **Key Optimizations:**
//  1. **Trie Construction**: O(W×L) where W is word count, L is average word length - done once upfront
//  2. **Efficient Word Finding**: Uses trie traversal instead of substring + map lookup
//  3. **Memory Efficient**: Reuses trie structure, minimal allocations during search
//  4. **Stable Sorting**: Primary sort by word count (descending), secondary by start position
//
// **Time Complexity:** O(N×M) where N is text length, M is maximum word length
// **Space Complexity:** O(W×L + S) where S is number of sequences found
//
// **How it works:**
//  1. Builds a trie from all dictionary words for fast prefix matching
//  2. For each position in text, greedily finds the longest sequence of consecutive words
//  3. Only keeps sequences with 2+ words
//  4. Sorts results by word count (descending) as required
//
// The algorithm handles the large input efficiently - 3M characters with 12K+ words should complete
// in reasonable time while finding all valid word sequences.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

const (
	kjvBaseFilename = "kjv-bag-of-words.txt"
	minPrintWordLen = 4
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	_, mainDir, _, _ := runtime.Caller(0)
	kjvFilename := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(mainDir))), kjvBaseFilename)
	buf, err := os.ReadFile(kjvFilename)
	must(err)

	words := map[string]struct{}{}
	for _, line := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		word := line[6:] // skip count before word, e.g. "63919 the"
		words[word] = struct{}{}
	}

	log.Printf("Loaded %v words", len(words))
	trie := buildTrie(words)

	for _, arg := range flag.Args() {
		processFile(arg, trie)
	}

	log.Printf("Done.")
}

func processFile(filename string, trie *trieNode) {
	buf, err := os.ReadFile(filename)
	must(err)

	ws := findWordSequences(string(buf), trie)
	if len(ws) == 0 {
		return
	}
	msg := fmt.Sprintf("%v: Found %v word sequences:", filename, len(ws))

	for _, seq := range ws {
		var hitMinPrintWordLen bool
		for _, w := range seq.Words {
			if len(w) >= minPrintWordLen {
				hitMinPrintWordLen = true
				break
			}
		}
		if hitMinPrintWordLen {
			if msg != "" {
				log.Printf("%v", msg)
				msg = "" // Only print this once per file
			}
			log.Printf("%v words: '%v'", seq.WordCount, strings.Join(seq.Words, " "))
		}
	}
}

type WordSequence struct {
	StartOffset int
	EndOffset   int
	WordCount   int
	Words       []string
}

func findWordSequences(text string, trie *trieNode) []*WordSequence {
	if len(text) == 0 {
		return nil
	}

	// Build trie for efficient prefix matching
	// trie := buildTrie(words)

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

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
