// -*- compile-command: "go run main.go"; -*-

// 01-breastplate attempts to reproduce the results of the first table
// in the book on page 18. It uses Numbers 4:11-20 with a skip of 22.
//
// There are a few minor differences, see comments below.
package main

import (
	"fmt"
	"log"
	"sync"

	codes "github.com/gmlewis/bible-codes"
	"github.com/gmlewis/bible-codes/strongs"
)

func main() {
	otRange, err := codes.NewOTRange("Numbers 4:11", "Numbers 4:20")
	must(err)

	table, err := otRange.GenTable(22, 48)
	must(err)

	fmt.Printf("table:\n%v\n\n", table)

	var wg sync.WaitGroup
	f := func(word string, entry *strongs.Entry) {
		defer wg.Done()
		english := entry.English()
		w, err := table.Find(word)
		if err != nil {
			// log.Printf("ERROR: %q: %v", english, err)
			return
		}
		if len(w) == 0 {
			// log.Printf("WARNING: no matches found for %q - %q", english, word)
			return
		}
		label := english
		if len(w) > 1 {
			label = fmt.Sprintf("%v (%v)", english, len(w))
		}
		fmt.Printf("Found: %q - %q\n", label, word)
	}

	for word, entry := range strongs.Hebrew {
		wg.Add(1)
		go f(word, entry)
	}

	wg.Wait()

	log.Printf("Done.")
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var words = map[string]string{
	// https://context.reverso.net/translation/english-hebrew/breastplate
	// not found: "The Breastplate (a)": "החושן",
	// not found: "The Breastplate (b)": "שריון החזה",
	// not found: "The Breastplate (c)": "חושן",
	// not found: "The Breastplate (d)": "החשן",
	// not found: "The Breastplate (e)": "המעטפת",
	// not found: "The Breastplate (f)": "היקר שלך",
	// "The Breastplate": "חשן", // "obtaining a judgement" according to: https://context.reverso.net/translation/hebrew-english/%D7%97%D7%A9%D7%9F
	"Obtaining a judgement": "חשן",
	// not found: "The Breastplate (h)": "שריון",

	// "Mediator": "מתווך", // https://context.reverso.net/translation/english-hebrew/mediator#%D7%9E%D7%AA%D7%95%D7%95%D7%9A
	// "Mediator": "מתווכ", // website could not find this word but is within the table

	// "Hidden, Mysterious": "תעלומה", // https://context.reverso.net/translation/english-hebrew/mysterious#%D7%AA%D7%A2%D7%9C%D7%95%D7%9E%D7%94
	// "Hidden, Mysterious": "עלומ", // website could not find this word but is within the table

	// https://context.reverso.net/translation/english-hebrew/circumcision#%D7%9E%D7%99%D7%9C%D7%94
	"Circumcision": "מילה",

	// https://context.reverso.net/translation/hebrew-english/%D7%94%D7%96%D7%94%D7%91
	"Gold": "הזהב",

	// "Shalom, Peace": "שלום", // https://context.reverso.net/translation/english-hebrew/shalom#%D7%A9%D7%9C%D7%95%D7%9D
	// "Shalom, Peace": "שלומ", // website could not find this word but is within the table

	// https://context.reverso.net/translation/english-hebrew/tabernacle#%D7%9E%D7%A9%D7%9B%D7%9F
	"Tabernacle": "משכן",

	// https://context.reverso.net/translation/english-hebrew/prophet#%D7%A0%D7%91%D7%99%D7%90
	// https://context.reverso.net/translation/hebrew-english/%D7%A0%D7%91%D7%90%D7%99
	// not found: "Prophet": "נביא", // Could not find same word as book for "Prophet", but instead, found:
	"Predict": "נבאי", // Which is the same Hebrew word that the book has for "Prophet", but has a different definition according to the website, which is also extremely cool.

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

	// https://context.reverso.net/translation/english-hebrew/Aaron#%D7%90%D7%94%D7%A8%D7%9F
	"Aaron": "אהרן",

	// https://context.reverso.net/translation/hebrew-english/%D7%94%D7%A7%D7%93%D7%A9
	"Consecration": "הקדש", // sanctuary, endowment, asylum, refuge, Holy

	// https://context.reverso.net/translation/english-hebrew/authentic#%D7%90%D7%9E%D7%99%D7%AA%D7%99
	"Authentic, True": "אמיתי",

	// https://context.reverso.net/translation/english-hebrew/faith#%D7%90%D7%9E%D7%95%D7%A0%D7%94
	"Faith": "אמונה",

	// "Trumpet Call": "תרועת",	// https://context.reverso.net/translation/hebrew-english/%D7%AA%D7%A8%D7%95%D7%A2%D7%AA
	"Fanfare": "תרועה", // https://context.reverso.net/translation/hebrew-english/%D7%AA%D7%A8%D7%95%D7%A2%D7%94
}
