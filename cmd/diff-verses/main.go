// -*- compile-command: "go run main.go sw1769_allwords_cleaned.txt > diffs.txt"; -*-

// diff-verses finds and prints the differences between two versions
// of the King James Bible.
//
// The first version is:
//
//	https://github.com/dewhisna/KingJamesPureBibleSearch/blob/master/text/build/sw1769/sw1769_allwords_cleaned.txt
//
// The second version is from: https://www.o-bible.com/download/kjv.txt
// which has been stored here:
//
//	https://github.com/gmlewis/bible-codes/blob/master/kjv-all-words.txt
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gmlewis/bible-codes/kjv"
)

const (
	label1 = "Pure Bible Search: "
	label2 = "Authorized 930105: "
)

func main() {
	flag.Parse()

	v1file := flag.Arg(0)
	buf, err := os.ReadFile(v1file)
	must(err)
	v1words := strings.Split(string(buf), "\n")
	log.Printf("Got %v words from file %v", len(v1words), v1file)

	v2, err := kjv.GetVerses()
	must(err)
	log.Printf("Got %v verses from kjv.txt", len(v2))

	printDiffs(v1words, v2)
}

func printDiffs(v1words []string, v2 []*kjv.Verse) {
	v1 := v1words[:]
	var totalWordDiffs int
	var totalVerseDiffs int
	for _, verse := range v2 {
		var wordDiffs int
		v1, wordDiffs = verseDiff(v1, verse)
		if wordDiffs > 0 {
			totalVerseDiffs++
			totalWordDiffs += wordDiffs
		}
	}
	fmt.Printf("\nFound %v total word differences and %v total verse differences\n",
		totalWordDiffs, totalVerseDiffs)
}

func verseDiff(v1 []string, verse *kjv.Verse) ([]string, int) {
	if v1[0] != verse.Words[0] {
		log.Fatalf("programming error: %q != %#v", v1[0], verse)
	}

	label2Indent := strings.Repeat(" ", len(label2))

	var wordDiffs int
	v1words := make([]string, 0, len(verse.Words))
	diffCarets := make([]string, 0, len(verse.Words))
	for _, w := range verse.Words {
		v1words = append(v1words, v1[0])
		w1 := v1[0]
		v1 = v1[1:]
		if w1 != w {
			wordDiffs++
			diffCarets = append(diffCarets, strings.Repeat("^", len(w)))
			continue
		}
		diffCarets = append(diffCarets, strings.Repeat(" ", len(w)))
	}

	if wordDiffs > 0 {
		var plural string
		if wordDiffs > 1 {
			plural = "s"
		}
		fmt.Printf("%v %v:%v has %v word difference%v:\n%v%v\n%v%v\n%v%v\n\n",
			verse.Book,
			verse.Chapter,
			verse.VerseNum,
			wordDiffs,
			plural,
			label1,
			strings.Join(v1words, " "),
			label2,
			strings.Join(verse.Words, " "),
			label2Indent,
			strings.Join(diffCarets, " "),
		)
	}

	return v1, wordDiffs
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
