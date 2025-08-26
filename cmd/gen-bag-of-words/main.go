// -*- compile-command: "go run main.go"; -*-

// gen-bag-of-words reads the "kjv-all-words.txt" file, removes all hyphens
// then lower-cases the words, finds their frequency, and prints them out
// with their corresponding usage counts in descending frequency order.
//
// For example:
//
//	go run main.go
//	Read a total of 789633 words
//	Found total of 12676 unique words
//	63919 the
//	...
//
// See kjv-bag-of-words.txt at base of repo.
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

const (
	kjvBaseFilename = "kjv-all-words.txt"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	_, mainDir, _, _ := runtime.Caller(0)
	kjvFilename := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(mainDir))), kjvBaseFilename)
	buf, err := os.ReadFile(kjvFilename)
	must(err)
	kjvText := strings.ReplaceAll(string(buf), "-", "")
	kjvText = strings.ReplaceAll(kjvText, "'", "")
	kjvText = strings.ToLower(strings.TrimSpace(kjvText))

	kjvWords := strings.Split(kjvText, "\n")

	log.Printf("Read a total of %v words", len(kjvWords))

	bag := map[string]int{}
	for _, word := range kjvWords {
		bag[word]++
	}

	log.Printf("Found total of %v unique words", len(bag))

	words := make([]string, 0, len(bag))
	for k := range bag {
		words = append(words, k)
	}

	sort.Slice(words, func(a, b int) bool {
		cmp := bag[words[b]] - bag[words[a]]
		if cmp == 0 {
			return words[a] < words[b]
		}
		return cmp < 0
	})

	for _, word := range words {
		log.Printf("%5d %v", bag[word], word)
	}

	log.Printf("Done.")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
