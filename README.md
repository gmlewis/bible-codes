# bible-codes

This is an attempt to first reproduce the findings made in this very interesting book:
https://www.amazon.com/gp/product/B09M83QHD5

The Hebrew and Greek Bible text and data is provided by this amazing website:
https://livinggreeknt.org/

The English Bible text has been proven by https://www.youtube.com/@TruthisChrist
(and others) to be the perfect Word of God in the English language only
in the King James Bible translation (which is in the public domain). The text
file `kjv.txt` was provided by this repo: https://www.o-bible.com/

## Processing

The file "raw-ot.txt" was created using the command:

```bash
cut -f8 LGNT-OT-Data.txt | tail -n +2 > raw-ot.txt
```

## Summary

Unfortunately, I was unable to duplicate the findings in the book
as you can also verify by running:

```bash
$ go run cmd/01-breastplate/main.go
...
$ go run cmd/02-holy-of-holies/main.go
...
```

However, during the course of my investigations, I came across another
website that strongly resonated with me: https://www.rapturepuzzle.com/

In case that website becomes unavailable, and with permission from
Renee Moses (the author), I wanted to make her PDF
files available for all to read here in this repo:

* [RAPTURE PUZZLE (May 28, 2023)](RAPTURE%20PUZZLE%20(May%2028%2C%202023).pdf)
* [Rapture Puzzle Summary](Rapture%20Puzzle%20Summary.pdf)
* [This Is My Story](This%20Is%20My%20Story.pdf)
