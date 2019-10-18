package currency

import "strconv"

// Currency is a decimal with precision of two * 100.
type Currency int

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

// func main() {
// }
