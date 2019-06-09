package unpack

import (
	"strconv"
	"strings"
	"unicode"
)

func Do(in string) string {
	escaped = false
	escapedNum = false
	doubleEscaped = false
	lastRune = '\\'

	var res strings.Builder
	for i, r := range in {
		res.WriteString(parse(i, r))
	}
	return res.String()
}

var (
	escaped       bool
	escapedNum    bool
	doubleEscaped bool
	lastRune      rune
)

func parse(i int, r rune) string {
	defer func() {
		escapedNum = escaped && unicode.IsDigit(r)
		escaped = r == '\\'
		doubleEscaped = escaped && lastRune == '\\'
		lastRune = r
	}()

	if unicode.IsDigit(r) && i == 0 {
		return ""
	}

	if r == '\\' && !escaped {
		return ""
	}

	if unicode.IsDigit(r) && unicode.IsDigit(lastRune) && !escaped && !escapedNum {
		return ""
	}

	if unicode.IsDigit(r) && !escaped ||
		unicode.IsDigit(r) && escapedNum ||
		unicode.IsDigit(r) && doubleEscaped {
		repeat, _ := strconv.Atoi(string(r))
		var res strings.Builder
		for repeat > 1 {
			res.WriteRune(lastRune)
			repeat--
		}
		return res.String()
	}

	return string(r)
}
