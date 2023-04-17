// Package strongs provides a map of all words in Strong's Concordance
// according to: https://www.kingjamesbibleonline.org/strongs-concordance/
package strongs

// Entry represents an entry in Strong's Concordance.
type Entry struct {
	Num  int
	Word string // original
	Desc string // HTML escaped string
}
