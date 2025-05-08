// -*- compile-command: "go install ."; -*-
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	log.SetFlags(0)
	flag.Parse()

	c := newClient()

	for _, arg := range flag.Args() {
		buf, err := os.ReadFile(arg)
		if err == nil {
			lines := strings.Split(string(buf), "\n")
			for _, line := range lines {
				if line == "" || strings.HasPrefix(line, "#!") || strings.HasPrefix(line, "exec ") {
					continue
				}
				c.processLine(line)
			}
			continue
		}
		c.processLine(arg)
	}

	log.Printf("Done.")
}

type clientT struct {
	kjv string
}

func newClient() *clientT {
	_, srcfile, _, _ := runtime.Caller(0)
	packageDir := filepath.Dir(srcfile)
	fullFilename := filepath.Join(filepath.Dir(packageDir), "../kjv.txt")
	buf, err := os.ReadFile(fullFilename)
	if err != nil {
		log.Fatal(err)
	}
	return &clientT{kjv: string(buf)}
}

func (c *clientT) processLine(line string) {
	parts := strings.Split(line, "-")
	if len(parts) == 1 {
		c.processVerse(line)
		return
	}
	colonIndex := strings.Index(parts[0], ":")
	if colonIndex < 0 {
		log.Printf("unable to parse: %v", line)
		return
	}
	firstVerseStr := parts[0][colonIndex+1:]
	firstVerse, err := strconv.Atoi(firstVerseStr)
	if err != nil {
		log.Printf("unable to parse firstVerse: %v", line)
		return
	}
	lastVerseStr := parts[1]
	lastVerse, err := strconv.Atoi(lastVerseStr)
	if err != nil {
		log.Printf("unable to parse lastVerse: %v", line)
		return
	}
	for verse := firstVerse; verse <= lastVerse; verse++ {
		v := fmt.Sprintf("%v:%v", parts[0][:colonIndex], verse)
		c.processVerse(v)
	}
}

func (c *clientT) processVerse(line string) {
	re := regexp.MustCompile("(?sm)^(" + line + ".*?)$")
	m := re.FindStringSubmatch(c.kjv)
	if len(m) > 0 {
		fmt.Printf("%v\n", m[0])
		return
	}
}
