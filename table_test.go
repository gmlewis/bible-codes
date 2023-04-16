package codes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

const (
	holyOfHolies   = "קדש הקדשים"
	theBreastplate = "החושן"
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
	tests := []struct {
		name string
		word string
		want []*Match
	}{
		{
			name: "Holy of Holies",
			word: holyOfHolies,
			want: []*Match{{RuneKeys: []Key{
				{18, 21}, {19, 21}, {20, 21}, {21, 21},
				{0, 22}, {1, 22}, {2, 22}, {3, 22}, {4, 22},
			}}},
		},
		{
			name: "TheBreastplate",
			word: theBreastplate,
			want: []*Match{{RuneKeys: []Key{
				{21, 7}, {21, 8}, {21, 9}, {21, 10}, {21, 11},
			}}},
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := table.Find(tt.word)
			if err != nil {
				t.Fatal(err)
			}

			if len(got) != len(tt.want) {
				t.Fatalf("table.Find found %v matches, want %v", len(got), len(tt.want))
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("table.Find mismatch (-want +got):\n%v", diff)
			}
		})
	}
}

func TestFewestDeltaPairs(t *testing.T) {
	table := getTable(t)
	tests := []struct {
		name string
		word string
		want int
	}{
		{
			name: "Holy of Holies",
			word: holyOfHolies,
			want: 271,
		},
		{
			name: "TheBreastplate",
			word: theBreastplate,
			want: 329,
		},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wordRunes := stripWhitespace(tt.word)
			deltas, err := table.fewestDeltaPairs(wordRunes)
			if err != nil {
				t.Fatal(err)
			}

			if got := len(deltas); got != tt.want {
				t.Errorf("fewestDeltaPairs(%q) = %v, want %v; deltas:\n%v", tt.word, got, tt.want, deltas)
			}
		})
	}
}

func TestPairDeltas(t *testing.T) {
	locs1 := keyPositionMapT{Key{0, 16}: struct{}{}, Key{0, 20}: struct{}{}, Key{2, 7}: struct{}{}, Key{2, 12}: struct{}{}, Key{2, 15}: struct{}{}, Key{4, 6}: struct{}{}, Key{4, 11}: struct{}{}, Key{5, 13}: struct{}{}, Key{5, 18}: struct{}{}, Key{6, 10}: struct{}{}, Key{6, 22}: struct{}{}, Key{7, 3}: struct{}{}, Key{7, 18}: struct{}{}, Key{7, 20}: struct{}{}, Key{8, 6}: struct{}{}, Key{8, 11}: struct{}{}, Key{9, 16}: struct{}{}, Key{10, 7}: struct{}{}, Key{12, 9}: struct{}{}, Key{12, 15}: struct{}{}, Key{12, 18}: struct{}{}, Key{13, 13}: struct{}{}, Key{13, 16}: struct{}{}, Key{14, 9}: struct{}{}, Key{14, 14}: struct{}{}, Key{15, 0}: struct{}{}, Key{15, 2}: struct{}{}, Key{15, 6}: struct{}{}, Key{17, 12}: struct{}{}, Key{17, 13}: struct{}{}, Key{17, 14}: struct{}{}, Key{17, 18}: struct{}{}, Key{18, 3}: struct{}{}, Key{18, 10}: struct{}{}, Key{19, 14}: struct{}{}, Key{20, 16}: struct{}{}, Key{20, 19}: struct{}{}, Key{20, 20}: struct{}{}, Key{20, 24}: struct{}{}, Key{21, 7}: struct{}{}, Key{21, 21}: struct{}{}}
	locs2 := keyPositionMapT{Key{0, 4}: struct{}{}, Key{1, 21}: struct{}{}, Key{3, 8}: struct{}{}, Key{6, 0}: struct{}{}, Key{6, 11}: struct{}{}, Key{10, 6}: struct{}{}, Key{11, 11}: struct{}{}, Key{12, 16}: struct{}{}, Key{18, 19}: struct{}{}, Key{20, 2}: struct{}{}, Key{20, 15}: struct{}{}, Key{21, 8}: struct{}{}}

	got := pairDeltas(locs1, locs2)
	if want := 412; len(got) != want {
		t.Fatalf("pairDeltas = %v deltas, want %v", len(got), want)
	}

	want := Key{0, 1}
	if _, ok := got[want]; !ok {
		t.Errorf("pairDeltas key %v is missing", want)
	}
}
