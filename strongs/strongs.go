// Package strongs provides a map of all words in Strong's Concordance
// according to: https://www.kingjamesbibleonline.org/strongs-concordance/
package strongs

import (
	"html"
	"strings"
)

// Entry represents an entry in Strong's Concordance.
type Entry struct {
	Num    int
	Word   string // original
	Length int    // number of runes in key value
	Desc   string // HTML escaped string
}

// English returns the KJV English usage of the Strong's entry.
func (e *Entry) English() string {
	htmlStr := html.UnescapeString(e.Desc)
	i := strings.Index(htmlStr, "KJV: ")
	if i < 0 {
		i = strings.Index(htmlStr, ":")
		if i >= 0 {
			htmlStr = htmlStr[:i]
		}
		return htmlStr
	}

	htmlStr = htmlStr[i+5:]
	j := strings.Index(htmlStr, ".")
	if j >= 0 {
		htmlStr = htmlStr[:j]
	}
	return htmlStr
}

// FilterByLength returns a smaller map based on filtering by a minimum word length.
func FilterByLength(orig map[string]*Entry, minLength int) map[string]*Entry {
	if minLength < 1 {
		return orig
	}
	result := map[string]*Entry{}
	for k, v := range orig {
		if v.Length < minLength {
			continue
		}
		result[k] = v
	}
	return result
}
