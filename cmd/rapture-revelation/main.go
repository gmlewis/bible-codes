// -*- compile-command: "go run main.go"; -*-

// rapture-revelation searches for any series of letters in any book (or all books)
// of the KJV using different skip codes and reports findings.
package main

import (
	"flag"
	"log"
	"strings"
	"sync"

	"github.com/gmlewis/advent-of-code-2021/enum"
	"github.com/gmlewis/bible-codes/kjv"
)

var (
	startOffset = flag.Int("offset", 0, "Starting offset to use")
	startSkip   = flag.Int("skip", 2, "Starting skip to use")

	lookFor    = flag.String("lookfor", "firstfruitsrapture", "Sequence of letters to look for (lower case)")
	maxWorkers = flag.Int("maxworkers", 1000, "Max number of concurrent goroutine workes to use")
	removePunc = flag.Bool("removepunc", false, "Remove apostrophes and dashes")
	searchBook = flag.String("search", "Revelation", "Which book to search ('all' for all books)")
)

func main() {
	flag.Parse()

	verses, err := kjv.GetVerses()
	must(err)
	log.Printf("Got %v verses", len(verses))

	if *searchBook == "all" {
		var fewestRunes int
		processedBooks := map[string]bool{}
		for _, verse := range verses {
			book := verse.Book
			if processedBooks[book] {
				continue
			}
			processedBooks[book] = true
			log.Printf("Searching book %q (%v of 66)...", book, len(processedBooks))
			numRunes := process(book, verses)
			if fewestRunes == 0 || numRunes < fewestRunes {
				fewestRunes = numRunes
			}
		}

		words := enum.FlatMap(verses, verse2words)
		log.Printf("Got %v words from all books", len(words))
		runes := words2runes(words)
		log.Printf("Got %v runes from all books", len(runes))
		processRunes(runes, fewestRunes/len(*lookFor))
	} else {
		process(*searchBook, verses)
	}

	log.Printf("Done.")
}

func words2runes(words []string) []rune {
	s := strings.Join(words, "")
	if *removePunc {
		s = strings.Replace(s, "'", "", -1)
		s = strings.Replace(s, "-", "", -1)
	}
	s = strings.ToLower(s)
	return []rune(s)
}

func process(book string, verses []*kjv.Verse) int {
	verses = enum.Filter(verses, filterBook(book))
	log.Printf("Got %v verses from the book of %q", len(verses), book)
	words := enum.FlatMap(verses, verse2words)
	log.Printf("Got %v words from the book of %q", len(words), book)
	runes := words2runes(words)
	log.Printf("Got %v runes from the book of %q", len(runes), book)

	processRunes([]rune(runes), *startSkip)

	return len(runes)
}

func processRunes(runes []rune, startSkip int) {
	reversed := make([]rune, len(runes))
	for i, r := range runes {
		reversed[len(runes)-i-1] = r
	}

	ch := make(chan struct{}, *maxWorkers)
	var wg sync.WaitGroup
	for skip := startSkip; skip < len(runes)/len(*lookFor); skip++ {
		searchWithSkip(runes, skip, false, wg, ch)
		searchWithSkip(reversed, skip, true, wg, ch)
	}

	wg.Wait()
}

func searchWithSkip(runes []rune, skip int, reversed bool, wg sync.WaitGroup, ch chan struct{}) {
	for offset := *startOffset; offset < skip; offset++ {
		wg.Add(1)
		ch <- struct{}{}
		go func(offset int) {
			searchWithOffsetAndSkip(runes, offset, skip, reversed)
			<-ch
			wg.Done()
		}(offset)
	}
}

func searchWithOffsetAndSkip(runes []rune, offset, skip int, reversed bool) {
	puz := make([]rune, 0, 1+len(runes)/skip)
	for i := offset; i < len(runes); i += skip {
		puz = append(puz, runes[i])
	}
	count := strings.Count(string(puz), *lookFor)
	if count == 0 {
		return
	}
	var isReversed string
	if reversed {
		isReversed = " (reversed)"
	}
	log.Printf("offset=%v, skip=%v, count=%v%v:\n%v", offset, skip, count, isReversed, string(puz))
}

func verse2words(verse *kjv.Verse) []string { return verse.Words }

func filterBook(book string) enum.FilterFunc[*kjv.Verse] {
	return func(verse *kjv.Verse) bool { return verse.Book == book }
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
