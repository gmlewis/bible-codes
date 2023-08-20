// Package kjv processes words and verses from "kjv.txt".
package kjv

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	filename = "kjv.txt"
)

// Verse presents a verse from the KJV.
type Verse struct {
	Book     string
	Chapter  int
	VerseNum int
	Words    []string
}

// GetVerses parses the verses from "kjv.txt".
func GetVerses() ([]*Verse, error) {
	_, srcfile, _, _ := runtime.Caller(0)
	packageDir := filepath.Dir(srcfile)

	fullFilename := filepath.Join(filepath.Dir(packageDir), filename)
	buf, err := os.ReadFile(fullFilename)
	if err != nil {
		return nil, err
	}

	verseLines := strings.Split(string(buf), "\n")[1:]

	verses := make([]*Verse, 0, len(verseLines)-1)
	for _, fullLine := range verseLines {
		fullLine = strings.TrimSpace(fullLine)
		if fullLine == "" {
			continue
		}
		words := strings.Split(fullLine, " ")
		verse := parseBCV(words[0])
		for _, word := range words[1:] {
			word = strings.Trim(word, ".,:;!?()")
			if word == "" {
				continue
			}
			verse.Words = append(verse.Words, word)
		}
		verses = append(verses, verse)
	}

	return verses, nil
}

func parseBCV(s string) *Verse {
	m := verseRE.FindStringSubmatch(s)
	if len(m) != 4 {
		log.Fatalf("unable to parse verse: %q", s)
	}
	book, ok := bookLookup[m[1]]
	if !ok {
		log.Fatalf("unable to parser verse book: %q", s)
	}
	chapter, err := strconv.Atoi(m[2])
	if err != nil {
		log.Fatalf("unable to parse verse chapter: %q", s)
	}
	verseNum, err := strconv.Atoi(m[3])
	if err != nil {
		log.Fatalf("unable to parse verse number: %q", s)
	}
	return &Verse{
		Book:     book,
		Chapter:  chapter,
		VerseNum: verseNum,
	}
}

var verseRE = regexp.MustCompile(`^(.*?)(\d+):(\d+)$`)

var bookLookup = map[string]string{
	"Ge":    "Genesis",
	"Exo":   "Exodus",
	"Lev":   "Leviticus",
	"Num":   "Numbers",
	"Deu":   "Deuteronomy",
	"Josh":  "Joshua",
	"Jdgs":  "Judges",
	"Ruth":  "Ruth",
	"1Sm":   "1 Samuel",
	"2Sm":   "2 Samuel",
	"1Ki":   "1 Kings",
	"2Ki":   "2 Kings",
	"1Chr":  "1 Chronicles",
	"2Chr":  "2 Chronicles",
	"Ezra":  "Ezra",
	"Neh":   "Nehemiah",
	"Est":   "Esther",
	"Job":   "Job",
	"Psa":   "Psalms",
	"Prv":   "Proverbs",
	"Eccl":  "Ecclesiastes",
	"SSol":  "Song of Solomon",
	"Isa":   "Isaiah",
	"Jer":   "Jeremiah",
	"Lam":   "Lamentations",
	"Eze":   "Ezekiel",
	"Dan":   "Daniel",
	"Hos":   "Hosea",
	"Joel":  "Joel",
	"Amos":  "Amos",
	"Obad":  "Obadiah",
	"Jonah": "Jonah",
	"Mic":   "Micah",
	"Nahum": "Nahum",
	"Hab":   "Habakkuk",
	"Zep":   "Zephaniah",
	"Hag":   "Haggai",
	"Zec":   "Zechariah",
	"Mal":   "Malachi",
	"Mat":   "Matthew",
	"Mark":  "Mark",
	"Luke":  "Luke",
	"John":  "John",
	"Acts":  "Acts",
	"Rom":   "Romams",
	"1Cor":  "1 Corinthians",
	"2Cor":  "2 Corinthians",
	"Gal":   "Galatians",
	"Eph":   "Ephesians",
	"Phi":   "Philippians",
	"Col":   "Colossians",
	"1Th":   "1 Thessalonians",
	"2Th":   "2 Thessalonians",
	"1Tim":  "1 Timothy",
	"2Tim":  "2 Timothy",
	"Titus": "Titus",
	"Phmn":  "Philemon",
	"Heb":   "Hebrews",
	"Jas":   "James",
	"1Pet":  "1 Peter",
	"2Pet":  "2 Peter",
	"1Jn":   "1 John",
	"2Jn":   "2 John",
	"3Jn":   "3 John",
	"Jude":  "Jude",
	"Rev":   "Revelation",
}
