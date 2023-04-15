# bible-codes

This is an attempt to first reproduce the findings made in this amazing book:
https://www.amazon.com/gp/product/B09M83QHD5

The Bible text and data is provided by this equally amazing website:
https://livinggreeknt.org/

## Processing

The file "raw-ot.txt" was created using the command:

```bash
cut -f8 LGNT-OT-Data.txt | tail -n +2 > raw-ot.txt
```
