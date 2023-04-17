// Package strongs provides a map of all words in Strong's Concordance
// according to: https://www.kingjamesbibleonline.org/strongs-concordance/
package strongs

import (
	"html"
	"strings"
)

// Entry represents an entry in Strong's Concordance.
type Entry struct {
	Num  int
	Word string // original
	Desc string // HTML escaped string
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
