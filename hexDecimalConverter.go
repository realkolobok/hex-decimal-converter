package hexDecConv

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func hexToDec(hexString string) (string, error) {
	var hexMap map[rune]int = map[rune]int{'a': 10, 'b': 11, 'c': 12, 'd': 13, 'e': 14, 'f': 15}
	hexFields := strings.Fields(hexString)
	var decimal []string
	for _, hex := range hexFields {
		var value int
		for i, h := range hex {
			var digit int
			if d, err := strconv.Atoi(string(h)); err == nil {
				digit = d
			} else if d, ok := hexMap[unicode.ToLower(h)]; ok {
				digit = d
			} else {
				return "", fmt.Errorf("Wrong input")
			}
			power := len(hex) - 1 - i
			value += digit * int(math.Pow(16, float64(power)))
		}
		decimal = append(decimal, strconv.Itoa(value))
	}
	return strings.Join(decimal, " "), nil
}

func decToHex(decString string) (string, error) {
	var hexMap map[int]string = map[int]string{10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f"}
	decFields := strings.Fields(decString)
	hex := []string{}
	for _, v := range decFields {
		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("You need to enter only numbers")
			return "", err
		}
		if val == 0 {
			hex = append(hex, "0")
			continue
		}
		var hexDigits []string
		quotient := val
		for quotient > 0 {
			remainder := quotient % 16
			quotient = quotient / 16

			if remainder >= 10 {
				hexDigits = append(hexDigits, hexMap[remainder])

			} else {
				hexDigits = append(hexDigits, strconv.Itoa(remainder))
			}
		}
		for i, j := 0, len(hexDigits)-1; i < j; i, j = i+1, j-1 {
			hexDigits[i], hexDigits[j] = hexDigits[j], hexDigits[i]
		}
		hex = append(hex, strings.Join(hexDigits, ""))

	}
	return strings.ToUpper(strings.Join(hex, " ")), nil
}
