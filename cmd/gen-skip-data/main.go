// -*- compile-command: "go run main.go -skip 777 -dir skip-777"; -*-

// gen-skip-data reads the kjv-all-letters.txt file and generates all
// possible offsets for a certain skip amount and writes the results
// to a named directory.
//
// For example:
//
//	gen-skip-data -skip 777 -dir skip-777
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

const (
	kjvBaseFilename = "kjv-all-letters.txt"
)

var (
	skipN   = flag.Int("skip", 0, "Amount of letters to skip")
	skipDir = flag.String("dir", "", "Directory to place the data (will delete contents if it exists)")
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	if *skipN == 0 {
		log.Fatal("Must provide -skip")
	}
	if *skipDir == "" {
		*skipDir = fmt.Sprintf("skip-%04d", *skipN)
	}

	if _, err := os.Stat(*skipDir); err == nil {
		must(os.RemoveAll(*skipDir))
	}
	must(os.Mkdir(*skipDir, 0755))

	_, mainDir, _, _ := runtime.Caller(0)
	kjvFilename := filepath.Join(filepath.Dir(filepath.Dir(filepath.Dir(mainDir))), kjvBaseFilename)
	buf, err := os.ReadFile(kjvFilename)
	must(err)
	kjvText := string(buf)

	numDigits := len(fmt.Sprintf("%v", *skipN))
	fmtStr := fmt.Sprintf("%v/skip-%v-%%0%vd.txt", *skipDir, *skipN, numDigits)
	for offset := range *skipN {
		filename := fmt.Sprintf(fmtStr, offset)
		genSkips(kjvText, *skipN, offset, writeSkipToFile(filename))
	}

	log.Printf("Done.")
}

type SkipWriter interface {
	writeSkip(text string)
}

func genSkips(text string, skipNum, offset int, w SkipWriter) {
	skipRE := fmt.Sprintf("%v(.)%v", strings.Repeat(".", offset), strings.Repeat(".", skipNum-1-offset))
	re := regexp.MustCompile(skipRE)
	out := re.ReplaceAllString(text, "$1")
	w.writeSkip(out)
}

type skipWriterT struct {
	filename string
}

func writeSkipToFile(filename string) *skipWriterT {
	return &skipWriterT{filename: filename}
}

func (s *skipWriterT) writeSkip(text string) {
	must(os.WriteFile(s.filename, []byte(text), 0644))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
