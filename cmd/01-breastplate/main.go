// -*- compile-command: "cd ../.. && go run cmd/01-breastplate/main.go"; -*-

// 01-breastplate attempts to reproduce the results of the first table
// in the book on page 18. It uses Numbers 4:11-20 with a skip of 22.
package main

import (
	"fmt"
	"log"

	codes "github.com/gmlewis/bible-codes"
)

func main() {
	otRange, err := codes.NewOTRange("Numbers 4:11", "Numbers 4:20")
	must(err)

	table, err := otRange.GenTable(22, 48)
	must(err)

	fmt.Printf("table:\n%v", table)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
