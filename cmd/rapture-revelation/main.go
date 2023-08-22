// -*- compile-command: "go run main.go"; -*-

// rapture-revelation searches for the word "rapture" in the book
// of Revelation using different skip codes and reports findings.
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
	lookFor    = flag.String("lookfor", "rapture", "Word to look for")
	searchBook = flag.String("search", "Revelation", "Which book to search")
)

func main() {
	flag.Parse()

	verses, err := kjv.GetVerses()
	must(err)

	log.Printf("Got %v verses", len(verses))
	verses = enum.Filter(verses, filterBook(*searchBook))
	log.Printf("Got %v verses from the book of %q", len(verses), *searchBook)
	words := enum.FlatMap(verses, verse2words)
	log.Printf("Got %v words from the book of %q", len(words), *searchBook)
	runes := strings.ToLower(strings.Join(words, ""))
	log.Printf("Got %v runes from the book of %q", len(runes), *searchBook)

	var wg sync.WaitGroup
	for skip := 2; skip < len(runes)/len(*lookFor); skip++ {
		wg.Add(1)
		go func(skip int) {
			searchWithSkip([]rune(runes), skip)
			wg.Done()
		}(skip)
	}

	wg.Wait()
	log.Printf("Done.")
}

func searchWithSkip(runes []rune, skip int) {
	var wg sync.WaitGroup
	for offset := 0; offset < skip; offset++ {
		wg.Add(1)
		go func(offset int) {
			searchWithOffsetAndSkip(runes, offset, skip)
			wg.Done()
		}(offset)
	}

	wg.Wait()
}

func searchWithOffsetAndSkip(runes []rune, offset, skip int) {
	puz := make([]rune, 0, 1+len(runes)/skip)
	for i := offset; i < len(runes); i += skip {
		puz = append(puz, runes[i])
	}
	count := strings.Count(string(puz), *lookFor)
	if count == 0 {
		return
	}
	log.Printf("offset=%v, skip=%v, count=%v:\n%v", offset, skip, count, string(puz))
}

func verse2words(verse *kjv.Verse) []string {
	return verse.Words
}

func filterBook(book string) enum.FilterFunc[*kjv.Verse] {
	return func(verse *kjv.Verse) bool { return verse.Book == book }
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
