package main

import (
	"fmt"
	"github.com/realkolobok/hex-decimal-converter"
)

func main() {
	//Hex to Decimal
	dec, err := hexDecConv.HexToDec("1a ff 4c")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(dec) // 26 255 76
	}

	//Decimal to Hex
	hex, err := hexDecConv.DecToHex("26 255 76")
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(hex) // 1A FF 4C
	}

	//Error test
	_, err = hexDecConv.HexToDec("VB") // V and B are not hexdecimal numbers
	if err != nil {
		fmt.Println("Error: ", err) // Error: Wrong input
	}
}
