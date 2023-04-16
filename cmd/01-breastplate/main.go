// -*- compile-command: "go run main.go"; -*-

// 01-breastplate attempts to reproduce the results of the first table
// in the book on page 18. It uses Numbers 4:11-20 with a skip of 22.
package main

import (
	"log"

	codes "github.com/gmlewis/bible-codes"
)

func main() {
	otRange, err := codes.NewOTRange("Numbers 4:11", "Numbers 4:20")
	must(err)

	table, err := otRange.GenTable(22, 48)
	must(err)

	// fmt.Printf("table:\n%v", table)

	for english, word := range words {
		w, err := table.Find(word)
		if err != nil {
			log.Printf("%q: %v", english, err)
			continue
		}
		log.Printf("Found %q - %q: %v", english, word, w)
	}
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var words = map[string]string{
	// https://context.reverso.net/translation/english-hebrew/breastplate#%D7%94%D7%97%D7%95%D7%A9%D7%9F
	"The Breastplate": "החושן",
	// https://context.reverso.net/translation/english-hebrew/mediator#%D7%9E%D7%AA%D7%95%D7%95%D7%9A
	"Mediator": "מתווך",
	// https://context.reverso.net/translation/english-hebrew/mysterious#%D7%AA%D7%A2%D7%9C%D7%95%D7%9E%D7%94
	"Hidden, Mysterious": "תעלומה", // Could not find same word as book.
	// https://context.reverso.net/translation/english-hebrew/circumcision#%D7%9E%D7%99%D7%9C%D7%94
	"Circumcision": "מילה",
	// https://context.reverso.net/translation/hebrew-english/%D7%94%D7%96%D7%94%D7%91
	"Gold": "הזהב",
	// https://context.reverso.net/translation/english-hebrew/shalom#%D7%A9%D7%9C%D7%95%D7%9D
	"Shalom, Peace": "שלום",
	// https://context.reverso.net/translation/english-hebrew/tabernacle#%D7%9E%D7%A9%D7%9B%D7%9F
	"Tabernacle": "משכן",
	// https://context.reverso.net/translation/english-hebrew/prophet#%D7%A0%D7%91%D7%99%D7%90
	"Prophet": "נביא", // Could not find same word as book for "Prophet", but instead, found:
	// https://context.reverso.net/translation/hebrew-english/%D7%A0%D7%91%D7%90%D7%99
	"Predict": "נבאי", // Which is the same Hebrew word as the book, but different definition, also extremely cool.
	// https://context.reverso.net/translation/english-hebrew/Holy+of+Holies#%D7%A7%D7%93%D7%A9+%D7%94%D7%A7%D7%93%D7%A9%D7%99%D7%9D
	"Holy of Holies": "קדש הקדשים",
	// https://context.reverso.net/translation/english-hebrew/Yahusha
	"Yahusha": "יהושע",
	// https://context.reverso.net/translation/english-hebrew/onyx#%D7%A9%D7%95%D7%94%D7%9D
	"Onyx": "שוהם",
	// https://context.reverso.net/translation/english-hebrew/ephod#%D7%90%D7%A4%D7%95%D7%93
	"Ephod": "אפוד",
	// https://context.reverso.net/translation/english-hebrew/stones#%D7%90%D7%91%D7%A0%D7%99
	"Stones": "אבני",
	// https://context.reverso.net/translation/english-hebrew/eleazar#%D7%90%D7%9C%D7%A2%D7%96%D7%A8
	"Eleazar": "אלעזר",
	// https://context.reverso.net/translation/english-hebrew/aaron#%D7%90%D7%A8%D7%95%D7%9F
	"Aaron (2)": "ארון",
	// https://context.reverso.net/translation/hebrew-english/%D7%94%D7%A7%D7%93%D7%A9
	"Consecration": "הקדש", // sanctuary, endowment, asylum, refuge, Holy
	// https://context.reverso.net/translation/english-hebrew/authentic#%D7%90%D7%9E%D7%99%D7%AA%D7%99
	"Authentic, True": "אמיתי",
	// https://context.reverso.net/translation/english-hebrew/faith#%D7%90%D7%9E%D7%95%D7%A0%D7%94
	"Faith": "אמונה",
	// https://context.reverso.net/translation/hebrew-english/%D7%AA%D7%A8%D7%95%D7%A2%D7%AA
	"Trumpet Call": "תרועת",
}
