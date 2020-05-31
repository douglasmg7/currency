package currency

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// Currency is a decimal with precision of two * 100.
type Currency int

// Return a formated currency to "1000000.00".
func (c Currency) ToString() string {
	str := strconv.Itoa(int(c))
	for len(str) < 3 {
		str = "0" + str
	}
	strLen := len(str)
	currency := []rune{}
	for i, ch := range str {
		if i == (strLen - 2) {
			currency = append(currency, '.')
		}
		currency = append(currency, ch)
	}
	return string(currency)
}

// Return a formated currency to "R$ 1.000.000,00".
func (c Currency) Format() string {
	str := strconv.Itoa(int(c))
	for len(str) < 3 {
		str = "0" + str
	}
	strLen := len(str)
	currency := []rune{'R', '$', ' '}
	for i, ch := range str {
		if i == (strLen-11) && (strLen != 11) {
			currency = append(currency, '.')
		}
		if i == (strLen-8) && (strLen != 8) {
			currency = append(currency, '.')
		}
		if i == (strLen-5) && (strLen != 5) {
			currency = append(currency, '.')
		}
		if i == (strLen - 2) {
			currency = append(currency, ',')
		}
		currency = append(currency, ch)
	}
	return string(currency)
}

// Parse to currency.
// sep is a decimal separator, must be "", "," or ".".
func Parse(str string, sep string) (Currency, error) {
	switch sep {
	case ",":
		str = strings.Replace(str, ".", "", -1)
		str = strings.Replace(str, ",", ".", -1)
	case ".":
	case "":
	default:
		return 0, errors.New(`Invalid separator, must be "", "," or "."`)
	}
	str = strings.TrimSpace(str)
	val64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	// return int(math.Round(val64 * 100)), nil
	return Currency(math.Round(val64 * 100)), nil
}
