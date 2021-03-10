package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(3999))
}

func intToRoman(num int) string {
	romans := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	builder := strings.Builder{}

	for _, roman := range romans {
		for num >= roman.value {
			builder.WriteString(roman.digit)
			num -= roman.value
		}
	}

	return builder.String()
}
