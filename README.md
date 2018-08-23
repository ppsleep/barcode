# Barcode

[![GoDoc](https://godoc.org/github.com/ppsleep/barcode?status.svg)](https://godoc.org/github.com/ppsleep/barcode)

## Introduction

This is a barcode generation package for Golang

## Example

This is a simple example of creating a barcode in Code 128-A encoding format.

```go
package main

import (
    "github.com/ppsleep/barcode"
    "github.com/ppsleep/barcode/code128"
    "image/png"
    "os"
)

func main() {
    code, _ := code128.A("CODE 128-A")
    r := barcode.Encode(code, 2, 50)
    file, _ := os.Create("128A.png")
    defer file.Close()
    png.Encode(file, r)
}
```

