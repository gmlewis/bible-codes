# bible-codes

This is a failed attempt to first reproduce the findings made in this very interesting book:
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
website that strongly resonated with me: https://www.eyesupandopen.org/

In case that website becomes unavailable,
I wanted to make his PDF file available for all to read here in this repo:

* [The LITTLE BOOK of Revelation Chapter 10](TheLittleBook.pdf)
