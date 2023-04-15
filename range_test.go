package codes

import "testing"

func TestNewOTRange(t *testing.T) {
	got, err := NewOTRange("Numbers 4:11", "Numbers 4:20")
	if err != nil {
		t.Fatal(err)
	}

	want := &OTRange{StartWordPos: 50939, EndWordPos: 51100}
	if got.StartWordPos != want.StartWordPos || got.EndWordPos != want.EndWordPos {
		t.Errorf("NewOTRange = %#v, want %#v", got, want)
	}
}
