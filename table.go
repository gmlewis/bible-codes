package codes

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
)

// Key represents a (col,row)/(x,y) position in the table.
// Upper right is (0,0) and reads right to left.
type Key [2]int

// keyPositionMapT represents all key locations found for a
// particular rune.
type keyPositionMapT map[Key]struct{}

// lookupT represents the lookup table which is arranged
// first by runes, which then points to a map of all
// Keys found in the table for this rune.
type lookupT map[rune]keyPositionMapT

// Table represents a table like is shown in the book.
type Table struct {
	cols   int
	rows   int
	grid   map[Key]rune
	lookup lookupT
}

// GenTable generates a table using the given range, skip, and offset.
func (o *OTRange) GenTable(skip, offset int) (*Table, error) {
	words := strings.Split(rawOT, "\n")[o.StartWordPos-1 : o.EndWordPos]

	var count, col int
	t := &Table{
		cols:   skip,
		grid:   map[Key]rune{},
		lookup: lookupT{},
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
			if _, ok := t.lookup[r]; !ok {
				t.lookup[r] = keyPositionMapT{}
			}
			t.lookup[r][k] = struct{}{}

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

// Match represents a search match in a table.
type Match struct {
	Delta    Key
	RuneKeys []Key
}

func stripWhitespace(word string) []rune {
	var wordRunes []rune // strip whitespace
	for _, r := range word {
		if r == ' ' {
			continue
		}
		wordRunes = append(wordRunes, r)
	}
	return wordRunes
}

// Find searches for a word in the table and returns all matches.
func (t *Table) Find(word string) ([]*Match, error) {
	wordRunes := stripWhitespace(word)
	if len(wordRunes) < 3 {
		return nil, fmt.Errorf("%q has %v runes, need at least 3", wordRunes, len(wordRunes))
	}

	deltas, err := t.fewestDeltaPairs(wordRunes)
	if err != nil {
		return nil, err
	}

	log.Printf("%q: got %v deltas", word, len(deltas))

	var wg sync.WaitGroup
	resultCh := make(chan *Match, 10)
	for delta := range deltas {
		wg.Add(1)
		go func(delta Key) {
			defer wg.Done()
			t.findMatchesWithDelta(wordRunes, delta, resultCh)
		}(delta)
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var results []*Match
	for result := range resultCh {
		results = append(results, result)
	}

	return results, nil
}

func (t *Table) findMatchesWithDelta(wordRunes []rune, delta Key, resultCh chan<- *Match) {
}

func (t *Table) fewestDeltaPairs(wordRunes []rune) (map[Key]bool, error) {
	var lowestProduct int
	var lastLocations, locs1, locs2 keyPositionMapT
	for i, r := range wordRunes {
		keys, ok := t.lookup[r]
		if !ok {
			return nil, fmt.Errorf("%q not found - missing %q", wordRunes, r)
		}

		// log.Printf("fewestDeltaPairs: %q rune %q found at %v locations", wordRunes, r, len(keys))
		if i == 0 {
			locs1 = keys
			lastLocations = keys
			lowestProduct = math.MaxInt
			continue
		}

		product := len(keys) * len(lastLocations)
		if product < lowestProduct {
			// log.Printf("fewestDeltaPairs: %q rune %q : %v * %v = new lowestProduct: %v", wordRunes, r, len(keys), len(lastLocations), product)
			locs1, locs2 = lastLocations, keys
			lowestProduct = product
		}
		lastLocations = keys
	}

	result := pairDeltas(locs1, locs2)
	return result, nil
}

func pairDeltas(last, next keyPositionMapT) map[Key]bool {
	result := map[Key]bool{}
	for k1 := range last {
		for k2 := range next {
			delta := Key{k2[0] - k1[0], k2[1] - k1[1]}
			result[delta] = true
		}
	}
	return result
}
