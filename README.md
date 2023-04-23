# bible-codes

This is an attempt to first reproduce the findings made in this very interesting book:
https://www.amazon.com/gp/product/B09M83QHD5

The Bible text and data is provided by this amazing website:
https://livinggreeknt.org/

## Processing

The file "raw-ot.txt" was created using the command:

```bash
cut -f8 LGNT-OT-Data.txt | tail -n +2 > raw-ot.txt
```

## Summary

Unfortunately, I was unable to duplicate the findings in the above book
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

* [RAPTURE PUZZLE (May 28, 2023)](RAPTURE PUZZLE (May 28, 2023).pdf)
* [Rapture Puzzle Summary](Rapture Puzzle Summary.pdf)
* [This Is My Story](This Is My Story.pdf)
