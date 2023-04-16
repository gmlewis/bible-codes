package codes

import (
	"fmt"
	"log"
	"math"
	"strings"
)

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

// Find searches for a word in the table and returns all matches.
func (t *Table) Find(word string) (*Table, error) {
	deltas, err := t.fewestDeltaPairs(word)
	if err != nil {
		return nil, err
	}

	log.Printf("%q: got %v deltas", word, len(deltas))

	// for delta := range deltas {
	// }

	return nil, nil
}

func (t *Table) fewestDeltaPairs(word string) (map[Key]bool, error) {
	var lowestProduct, wordLen int
	var lastLocations, locs1, locs2 []Key
	for i, r := range word {
		if r == ' ' {
			if i == 0 {
				return nil, fmt.Errorf("first rune in %q cannot start with space", word)
			}
			continue
		}
		wordLen++
		keys, ok := t.lookup[r]
		if !ok {
			return nil, fmt.Errorf("%q not found - missing %q", word, r)
		}

		// log.Printf("fewestDeltaPairs: %q rune %q found at %v locations", word, r, len(keys))
		if i == 0 {
			locs1 = keys
			lastLocations = keys
			lowestProduct = math.MaxInt
			continue
		}

		product := len(keys) * len(lastLocations)
		if product < lowestProduct {
			// log.Printf("fewestDeltaPairs: %q rune %q : %v * %v = new lowestProduct: %v", word, r, len(keys), len(lastLocations), product)
			locs1, locs2 = lastLocations, keys
			lowestProduct = product
		}
		lastLocations = keys
	}

	if wordLen < 3 {
		return nil, fmt.Errorf("%q has %v runes, need at least 3", word, wordLen)
	}

	result := pairDeltas(locs1, locs2)
	return result, nil
}

func pairDeltas(last, next []Key) map[Key]bool {
	result := map[Key]bool{}
	for _, k1 := range last {
		for _, k2 := range next {
			delta := Key{k1[0] - k2[0], k1[1] - k2[1]}
			result[delta] = true
		}
	}
	return result
}
