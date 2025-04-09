# Hex-Decimal Converter (Go Library)

A small Go package for converting between hexadecimal and decimal numbers.

## 📦 Features

- Convert hex to decimal
- Convert decimal to hex
- Handles multi-number strings
- Validates input for errors

## 🚀 Installation
```bash
go get github.com/realkolobok/hex-decimal-converter
```
## 📚 Usage

```go
import "github.com/realkolobok/hex-decimal-converter"

dec, _ := hexDecConv.hexToDec("1A 2F")
// "26 47"

hex, _ := hexDecConv.decToHex("26 47")
// "1A 2F"
```

See full [examples](examples/) in the repo.
