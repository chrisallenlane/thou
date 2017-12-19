thou
====
`thou` inserts thousands-separators into `int`, `float64`, or `string`
representations of numbers. The thousands-separator, radix, and decimal
precision are configurable.


Installing
----------
```sh
go get github.com/chrisallenlane/thou
```

Usage
-----
```go
package main

import (
  "fmt"
  "github.com/chrisallenlane/thou"
)

func main() {
  // thousands separator, radix, and decimal precision
  sep  := ","
  rad  := "."
  prec := 4 

  // separate an integer
  out1 := thou.SepI(1000000, sep) 

  // separate a float
  out2, _ := thou.SepF(1000000.0000, prec, sep, rad)

  // separate a string
  out3, _ := thou.SepS("1000000.0000", prec, sep, rad)

  // display output
  fmt.Println(out1)
  fmt.Println(out2)
  fmt.Println(out3)

  /*

  Output:

  1,000,000
  1,000,000.0000
  1,000,000.0000

  */
}
```
