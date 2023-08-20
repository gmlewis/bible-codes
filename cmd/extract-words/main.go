// -*- compile-command: "go run main.go ../../kjv.txt > all-words.txt"; -*-

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	arg := flag.Arg(0)
	log.Printf("Reading file %v ...", arg)
	buf, err := os.ReadFile(arg)
	must(err)

	verses := strings.Split(string(buf), "\n")[1:]
	log.Printf("Got %v verses", len(verses))

	for _, fullLine := range verses {
		fullLine = strings.TrimSpace(fullLine)
		words := strings.Split(fullLine, " ")
		// vers := words[0]
		for _, word := range words[1:] {
			word = strings.Trim(word, ".,:;!?()")
			if word == "" {
				continue
			}
			fmt.Println(word)
		}
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
