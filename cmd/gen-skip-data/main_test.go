package main

import (
	"fmt"
	"testing"
)

func TestGenSkips(t *testing.T) {
	text := "abcdefghijklmnopqrstuvwxyz"

	tests := []struct {
		skipNum int
		offset  int
		want    string
	}{
		{
			skipNum: 7,
			offset:  0,
			want:    "ahov",
		},
		{
			skipNum: 7,
			offset:  1,
			want:    "bipw",
		},
		{
			skipNum: 7,
			offset:  2,
			want:    "cjqx",
		},
		{
			skipNum: 7,
			offset:  3,
			want:    "dkry",
		},
		{
			skipNum: 7,
			offset:  4,
			want:    "elsz",
		},
		{
			skipNum: 7,
			offset:  5,
			want:    "fmt",
		},
		{
			skipNum: 7,
			offset:  6,
			want:    "gnu",
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("skip %v, %v", tt.skipNum, tt.offset)
		t.Run(name, func(t *testing.T) {
			f := &fakeSkipWriter{}
			genSkips(text, tt.skipNum, tt.offset, f)

			if f.out != tt.want {
				t.Errorf("genSkips = %q,\nwant: %q", f.out, tt.want)
			}
		})
	}
}

type fakeSkipWriter struct {
	out string
}

func (f *fakeSkipWriter) writeSkip(text string) {
	f.out = text
}
