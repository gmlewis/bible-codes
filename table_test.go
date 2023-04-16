package codes

import (
	"testing"
)

const (
	holyOfHolies = "קודש הקודשים" // "Holy of Holies"
)

func getTable(t *testing.T) *Table {
	t.Helper()

	otRange, err := NewOTRange("Numbers 4:11", "Numbers 4:20")
	if err != nil {
		t.Fatal(err)
	}

	table, err := otRange.GenTable(22, 48)
	if err != nil {
		t.Fatal(err)
	}

	return table
}

func TestFind(t *testing.T) {
	table := getTable(t)

	got, err := table.Find(holyOfHolies)
	if err != nil {
		t.Fatal(err)
	}

	t.Errorf("table.Find = %v", got)
}

func TestFewestDeltaPairs(t *testing.T) {
	table := getTable(t)
	wordRunes := stripWhitespace(holyOfHolies)
	deltas, err := table.fewestDeltaPairs(wordRunes)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(deltas), 331; got != want {
		t.Errorf("fewestDeltaPairs(%q) = %v, want %v", holyOfHolies, got, want)
	}
}
