package codes

import (
	_ "embed"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	otWordPosCol = 4
)

//go:embed LGNT-OT-Data.txt
var otData string

//go:embed raw-ot.txt
var rawOT string

// OTRange represents an inclusive range of text from the Old Testament in Hebrew.
type OTRange struct {
	StartWordPos int // 1-indexed
	EndWordPos   int // 1-indexed
}

// NewOTRange parses start and end locations in the Old Testament and returns a new OTRange.
// Both start and end need 3 parts: Book, Chapter, and Verse.
// e.g. NewOTRange("Numbers 4:11", "Numbers 4:20")
// Note that Nehemiah 7:68 is a verse that is not in the Hebrew text.
func NewOTRange(start, end string) (*OTRange, error) {
	if start == "Nehemiah 7:68" || end == "Nehemiah 7:68" {
		return nil, errors.New("Nehemiah 7:68 is not in the Hebrew text")
	}

	startPos, err := getOTWordPos(start, false)
	if err != nil {
		return nil, err
	}

	endPos, err := getOTWordPos(end, true)
	if err != nil {
		return nil, err
	}

	return &OTRange{StartWordPos: startPos, EndWordPos: endPos}, nil
}

func getOTWordPos(loc string, findLast bool) (int, error) {
	parts := strings.Split(loc, " ")
	if len(parts) != 2 {
		return 0, fmt.Errorf("unable to parse %q", loc)
	}
	book := parts[0]
	next := strings.Split(parts[1], ":")
	if len(next) != 2 {
		return 0, fmt.Errorf("unable to parse %q", loc)
	}
	chapter := next[0]
	verse := next[1]

	find := strings.Index
	if findLast {
		find = strings.LastIndex
	}
	search := fmt.Sprintf("%v\t%v\t%v\t", book, chapter, verse)
	offset := find(otData, search)
	if offset < 0 {
		return 0, fmt.Errorf("%q not found", search)
	}

	line := otData[offset:]
	offset = strings.Index(line, "\n")
	line = line[:offset]
	cols := strings.Split(line, "\t")

	col, err := strconv.Atoi(cols[otWordPosCol])
	if err != nil {
		return 0, fmt.Errorf("error parsing %q", line)
	}

	return col, nil
}

// Key represents a (col,row)/(x,y) position in the table.
// Upper right is (0,0) and reads right to left.
type Key [2]int

// Table represents a table like is shown in the book.
type Table struct {
	cols   int
	rows   int
	grid   map[Key]rune
	lookup map[rune][]Key
}

// GenTable generates a table using the given range, skip, and offset.
func (o *OTRange) GenTable(skip, offset int) (*Table, error) {
	words := strings.Split(rawOT, "\n")[o.StartWordPos-1 : o.EndWordPos]

	var count, col int
	t := &Table{
		cols:   skip,
		grid:   map[Key]rune{},
		lookup: map[rune][]Key{},
	}
	for _, word := range words {
		// log.Printf("count=%v: word=%q", count, word)
		for _, r := range word {
			count++
			if count <= offset {
				continue
			}

			k := Key{col, t.rows}
			t.grid[k] = r
			t.lookup[r] = append(t.lookup[r], k)

			col++
			if col >= skip {
				col = 0
				t.rows++
			}
		}
	}
	log.Printf("GenTable(%v,%v): got %v words, %v runes, and %v grid runes", skip, offset, len(words), count, len(t.grid))

	return t, nil
}

// String draws a text representation of the Table.
func (t *Table) String() string {
	var lines []string
	for row := 0; row < t.rows; row++ {
		var line string
		for col := 0; col < t.cols; col++ {
			// k := Key{t.cols - col - 1, row} // reverse
			k := Key{col, row} // reverse
			line += string(t.grid[k])
		}
		lines = append(lines, line)
	}
	return strings.Join(lines, "\n")
}
