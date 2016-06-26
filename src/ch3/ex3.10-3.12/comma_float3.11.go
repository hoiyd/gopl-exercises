package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(commaFloat("+12.9"))
	fmt.Println(commaFloat("-123.99"))
	fmt.Println(commaFloat("+1234.9999"))
	fmt.Println(commaFloat("-12345.9999"))
	fmt.Println(commaFloat("+123456.99999"))
	fmt.Println(commaFloat("-1234567.999999"))
}

func commaFloat(value string) string {
	signal, s, suffix := preProcess(value)
	return signal + comma(s) + "." + suffix
}

func comma(s string) string {
	n := len(s)

	if n <= 3 {
		return s
	}

	r := n % 3
	var buf bytes.Buffer
	if r > 0 {
		buf.WriteString(s[:r])
		buf.WriteString(",")
	}

	for i := r; i < n; i = i + 3 {
		buf.WriteString(s[i : i+3])
		if i+3 < n {
			buf.WriteString(",")
		}
	}
	return buf.String()
}

func preProcess(s string) (string, string, string) {
	var signal string
	var suffix string
	if s[0] == '+' {
		signal = "+"
		s = s[1:]
	} else if s[0] == '-' {
		signal = "-"
		s = s[1:]
	} else {
		signal = ""
	}

	indexOfDecimal := strings.Index(s, ".")
	if indexOfDecimal < 0 {
		suffix = ""
	} else {
		suffix = s[indexOfDecimal+1:]
		s = s[:indexOfDecimal]
	}
	return signal, s, suffix
}
