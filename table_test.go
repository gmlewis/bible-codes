package codes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	holyOfHolies = "קדש הקדשים" // "Holy of Holies"
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

	if want := 1; len(got) != want {
		t.Fatalf("table.Find found %v matches, want %v", len(got), want)
	}

	want := []*Match{{RuneKeys: []Key{
		{18, 21}, {19, 21}, {20, 21}, {21, 21},
		{0, 22}, {1, 22}, {2, 22}, {3, 22}, {4, 22},
	}}}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("table.Find mismatch (-want +got):\n%v", diff)
	}
}

func TestFewestDeltaPairs(t *testing.T) {
	table := getTable(t)
	wordRunes := stripWhitespace(holyOfHolies)
	deltas, err := table.fewestDeltaPairs(wordRunes)
	if err != nil {
		t.Fatal(err)
	}

	if got, want := len(deltas), 271; got != want {
		t.Errorf("fewestDeltaPairs(%q) = %v, want %v", holyOfHolies, got, want)
	}
}
