// -*- compile-command: "cd ../.. && go run cmd/01-breastplate/main.go"; -*-

// 01-breastplate attempts to reproduce the results of the first table
// in the book on page 18. It uses Numbers 4:11-20 with a skip of 22.
package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	buf, err := os.ReadFile("raw-ot.txt")
	must(err)

	words := strings.Split(strings.TrimSpace(string(buf)), "\n")

	log.Printf("Got %v bytes, %v words", len(buf), len(words))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
