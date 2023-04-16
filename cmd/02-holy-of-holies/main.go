// -*- compile-command: "go run main.go"; -*-

// 02-holy-of-holies attempts to reproduce the results of the second table
// in the book on page 20. It uses Genesis 47:31-48:11 with a skip of 24.
//
// There are a few minor differences, see comments below.
package main

import (
	"fmt"
	"log"

	codes "github.com/gmlewis/bible-codes"
)

func main() {
	otRange, err := codes.NewOTRange("Genesis 47:31", "Genesis 48:11")
	must(err)

	table, err := otRange.GenTable(24, 13)
	must(err)

	fmt.Printf("table:\n%v\n\n", table)

	for english, word := range words {
		w, err := table.Find(word)
		if err != nil {
			log.Printf("ERROR: %q: %v", english, err)
			continue
		}
		if len(w) == 0 {
			log.Printf("WARNING: no matches found for %q - %q", english, word)
			continue
		}
		label := english
		if len(w) > 1 {
			label = fmt.Sprintf("%v (%v)", english, len(w))
		}
		fmt.Printf("Found: %q - %q\n", label, word)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var words = map[string]string{
	// https://context.reverso.net/translation/english-hebrew/Holy+of+Holies#%D7%A7%D7%93%D7%A9+%D7%94%D7%A7%D7%93%D7%A9%D7%99%D7%9D
	"Holiness":           "קדש",
	"The Holy of Holies": "הקדשים",

	// "The Breastplate": "חשכ", // website could not find this word but is within the table

	// https://context.reverso.net/translation/english-hebrew/commandment#%D7%9E%D7%A6%D7%95%D7%95%D7%94
	// "Commandment": "מצווה",
	// "Commandment": "םצווה", // website could not find this word but is within the table

	// "Ephod": "אפוד", // https://context.reverso.net/translation/english-hebrew/ephod#%D7%90%D7%A4%D7%95%D7%93
	// "Ephod": "אפד", // website could not find this word but is within the table

	// https://context.reverso.net/translation/english-hebrew/Uzziah#%D7%A2%D7%96%D7%99%D7%94
	"Uzziah": "עזיה",

	// https://context.reverso.net/translation/english-hebrew/temple#%D7%94%D7%99%D7%9B%D7%9C
	// "Temple, Sanctuary": "היכל",
	// "Temple, Sanctuary": "היךל", // website could not find this word but is within the table

	// https://context.reverso.net/translation/english-hebrew/bethlehem#%D7%91%D7%99%D7%AA+%D7%9C%D7%97%D7%9D
	"Bethlehem": "בית לחם",

	// https://context.reverso.net/translation/english-hebrew/covenant#%D7%91%D7%A8%D7%99%D7%AA
	"Covenant": "ברית",

	// https://context.reverso.net/translation/english-hebrew/utterance#%D7%90%D7%9E%D7%99%D7%A8%D7%94
	"Utterance": "אמירה",

	// https://context.reverso.net/translation/english-hebrew/priest#%D7%9B%D7%95%D7%9E%D7%A8
	// "Priest": "כומר",
	// "Priest": "כוהנ", // website could not find this word but is within the table

	// Successful matches from 01-breastplate:

	// https://context.reverso.net/translation/english-hebrew/breastplate
	// "The Breastplate": "חשכ", // website could not find this word but is within the table
	// not found: "The Breastplate (a)": "החושן",
	// not found: "The Breastplate (b)": "שריון החזה",
	// not found: "The Breastplate (c)": "חושן",
	// not found: "The Breastplate (d)": "החשן",
	// not found: "The Breastplate (e)": "המעטפת",
	// not found: "The Breastplate (f)": "היקר שלך",
	// "The Breastplate": "חשן", // "obtaining a judgement" according to: https://context.reverso.net/translation/hebrew-english/%D7%97%D7%A9%D7%9F
	"Obtaining a judgement": "חשן",
	// not found: "The Breastplate (h)": "שריון",

	// https://context.reverso.net/translation/english-hebrew/circumcision#%D7%9E%D7%99%D7%9C%D7%94
	"Circumcision": "מילה",

	// https://context.reverso.net/translation/hebrew-english/%D7%94%D7%96%D7%94%D7%91
	"Gold": "הזהב",

	// https://context.reverso.net/translation/english-hebrew/onyx#%D7%A9%D7%95%D7%94%D7%9D
	"Onyx": "שוהם",

	// https://context.reverso.net/translation/english-hebrew/stones#%D7%90%D7%91%D7%A0%D7%99
	"Stones": "אבני",

	// https://context.reverso.net/translation/hebrew-english/%D7%94%D7%A7%D7%93%D7%A9
	"Consecration": "הקדש", // sanctuary, endowment, asylum, refuge, Holy
}
